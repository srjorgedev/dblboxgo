package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/srjorgedev/dblboxgo/internal/domain/unit"
	"github.com/srjorgedev/dblboxgo/pkg"
)

type SQLUnitRepository struct {
	db        *sql.DB
	cache     []*unit.UnitSummary
	cacheLock sync.RWMutex
}

func NewSQLUnitRepository(db *sql.DB) *SQLUnitRepository {
	repo := &SQLUnitRepository{db: db}

	go func() {
		units, err := repo.GetAllUnitSummariesCached()
		if err != nil {
			fmt.Printf("[ API ] Failed to preload units: %v\n", err)
			return
		}
		repo.cacheLock.Lock()
		repo.cache = units
		repo.cacheLock.Unlock()
		fmt.Printf("[ API ] Preloaded %d units into cache\n", len(units))
	}()

	return repo
}

func (r *SQLUnitRepository) GetUnitByID(id string) (*unit.Unit, error) {
	// Obtener datos de la tabla principal
	general := &unit.UnitParametersGeneral{}
	err := r.db.QueryRow(`
		SELECT
			u._id, u.num_id, 
			c._id, c.data_es,
			r._id, r.data_es, 
			t._id, t.data_es, 
			u.transform, u.lf, u.zenkai, u.kit_type, u.zenkai_kit_type, u.tag_switch
		FROM unit_parameters_general u
		JOIN data_chapter c ON u.chapter = c._id
		JOIN data_rarity r ON u.rarity = r._id
		JOIN data_type t ON u.type = t._id
		WHERE u._id = ?
	`, id).Scan(
		&general.ID,
		&general.NumID,
		&general.Chapter.ID, &general.Chapter.DataEs,
		&general.Rarity.ID, &general.Rarity.DataEs,
		&general.Type.ID, &general.Type.DataEs,
		&general.Transform,
		&general.LF,
		&general.Zenkai,
		&general.KitType,
		&general.ZenkaiKitType,
		&general.TagSwitch,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit: %w", err)
	}

	// Crear la estructura Unit
	u := &unit.Unit{
		ID:            general.ID,
		NumID:         general.NumID,
		Chapter:       general.Chapter,
		Rarity:        general.Rarity,
		Type:          general.Type,
		Transform:     general.Transform,
		LF:            general.LF,
		Zenkai:        general.Zenkai,
		KitType:       general.KitType,
		ZenkaiKitType: general.ZenkaiKitType,
		TagSwitch:     general.TagSwitch,
	}

	// Obtener datos relacionados
	names, err := r.GetUnitNames(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit names: %w", err)
	}
	u.Names = names

	tags, err := r.GetUnitTags(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit tags: %w", err)
	}
	u.Tags = tags

	affinity, err := r.GetUnitAffinity(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit affinity: %w", err)
	}
	u.Affinity = affinity

	heldCards, err := r.GetUnitHeldCards(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit held cards: %w", err)
	}
	u.HeldCards = heldCards

	u.Images = pkg.GetUnitImages(u.NumID, u.Transform, u.TagSwitch)

	return u, nil
}

func (r *SQLUnitRepository) GetAllUnitSummariesCached() ([]*unit.UnitSummary, error) {
	fmt.Println("[ API ] Checking if cache exists")

	// Check if cache exists
	r.cacheLock.RLock()
	if r.cache != nil {
		defer r.cacheLock.RUnlock()
		return r.cache, nil
	}
	r.cacheLock.RUnlock()

	fmt.Println("[ API ] Cache not found, fetching from DB")

	// Cache not set, fetch from DB
	units, err := r.GetAllUnitSummaries()
	if err != nil {
		return nil, err
	}

	// Save to cache
	r.cacheLock.Lock()
	r.cache = units
	r.cacheLock.Unlock()

	return units, nil
}

func (r *SQLUnitRepository) GetAllUnitSummaries() ([]*unit.UnitSummary, error) {

	rows, err := r.db.Query(`
		SELECT 
			u._id,
			u.num_id,
			c._id     AS chapter_id,
			r._id     AS rarity_id,
			t._id     AS type_id,
			u.transform,
			u.lf,
			u.zenkai,
			u.tag_switch,
			GROUP_CONCAT(DISTINCT tg._id) AS tags,
			GROUP_CONCAT(DISTINCT af._id) AS affinity
		FROM unit_parameters_general u
		JOIN data_chapter c ON u.chapter = c._id
		JOIN data_rarity r  ON u.rarity  = r._id
		JOIN data_type t    ON u.type    = t._id
		LEFT JOIN unit_parameters_tag ut ON u._id = ut._unit_id
		LEFT JOIN data_tag tg ON ut.tag = tg._id
		LEFT JOIN unit_parameters_affinity ua ON u._id = ua._unit_id
		LEFT JOIN data_affinity af ON ua.affinity = af._id
		GROUP BY u._id
		ORDER BY u._id DESC;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get units: %w", err)
	}
	defer rows.Close()

	var units []*unit.UnitSummary

	for rows.Next() {
		u := &unit.UnitSummary{}

		var tagsRaw, affinityRaw sql.NullString

		if err := rows.Scan(
			&u.ID,
			&u.NumID,
			&u.Chapter,
			&u.Rarity,
			&u.Type,
			&u.Transform,
			&u.LF,
			&u.Zenkai,
			&u.TagSwitch,
			&tagsRaw,
			&affinityRaw,
		); err != nil {
			return nil, fmt.Errorf("failed to scan unit: %w", err)
		}

		u.Images = pkg.GetUnitImages(u.NumID, u.Transform, u.TagSwitch)

		// Parse tags into []int
		if tagsRaw.Valid && tagsRaw.String != "" {
			for _, t := range strings.Split(tagsRaw.String, ",") {
				if id, err := strconv.Atoi(t); err == nil {
					u.Tags = append(u.Tags, id)
				}
			}
		}

		// Parse affinity into []int
		if affinityRaw.Valid && affinityRaw.String != "" {
			for _, a := range strings.Split(affinityRaw.String, ",") {
				if id, err := strconv.Atoi(a); err == nil {
					u.Affinity = append(u.Affinity, id)
				}
			}
		}

		units = append(units, u)
	}

	return units, nil
}

// func (r *SQLUnitRepository) GetAllUnitTierMaker() ([]*unit.UnitTierMaker, error) {
// 	bchacutURL := os.Getenv("BCHACUT_URL")
// 	bchaicoURL := os.Getenv("BCHAICO_URL")

// 	// Query all general unit info
// 	rows, err := r.db.Query(`
//         SELECT
//             u._id, u.num_id,
//             c._id,
//             r._id,
//             t._id,
//             u.transform, u.lf, u.zenkai, u.tag_switch
//         FROM unit_parameters_general u
//         JOIN data_chapter c ON u.chapter = c._id
//         JOIN data_rarity r ON u.rarity = r._id
//         JOIN data_type t ON u.type = t._id
// 		ORDER BY u._id DESC
//     `)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get units: %w", err)
// 	}
// 	defer rows.Close()

// 	var units []*unit.UnitTierMaker

// 	for rows.Next() {
// 		u := &unit.UnitTierMaker{}
// 		if err := rows.Scan(
// 			&u.ID,
// 			&u.NumID,
// 			&u.Chapter,
// 			&u.Rarity,
// 			&u.Type,
// 			&u.Transform,
// 			&u.TagSwitch,
// 		); err != nil {
// 			return nil, fmt.Errorf("failed to scan unit: %w", err)
// 		}

// 		// Add images
// 		baseImage := unit.Images{
// 			BChaCut: bchacutURL + strconv.Itoa(u.NumID) + ".webp",
// 			BChaIco: bchaicoURL + strconv.Itoa(u.NumID) + ".webp",
// 		}
// 		u.Images = append(u.Images, baseImage)

// 		if u.Transform || u.TagSwitch {
// 			image2 := unit.Images{
// 				BChaCut: bchacutURL + strconv.Itoa(u.NumID) + "2.webp",
// 				BChaIco: bchaicoURL + strconv.Itoa(u.NumID) + "2.webp",
// 			}
// 			u.Images = append(u.Images, image2)
// 		}
// 		if u.Transform && u.TagSwitch {
// 			image3 := unit.Images{
// 				BChaCut: bchacutURL + strconv.Itoa(u.NumID) + "3.webp",
// 				BChaIco: bchaicoURL + strconv.Itoa(u.NumID) + "3.webp",
// 			}
// 			u.Images = append(u.Images, image3)
// 		}

// 		// Add tags
// 		tags, err := r.GetUnitTags(u.ID)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get unit tags: %w", err)
// 		}
// 		u.Tags = make([]int, len(tags))

// 		for i, t := range tags {
// 			u.Tags[i] = t.ID
// 		}

// 		// Add affinity
// 		affinity, err := r.GetUnitAffinity(u.ID)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get unit affinity: %w", err)
// 		}
// 		u.Affinity = make([]int, len(affinity))

// 		for i, a := range affinity {
// 			u.Affinity[i] = a.ID
// 		}

// 		units = append(units, u)
// 	}

// 	return units, nil
// }

func (r *SQLUnitRepository) GetUnitSummaryByID(id string) (*unit.UnitSummary, error) {
	u := &unit.UnitSummary{}
	err := r.db.QueryRow(`
		SELECT
			u._id, u.num_id, 
			c._id,
			r._id, 
			t._id, 
			u.transform, u.lf, u.zenkai, u.tag_switch
		FROM unit_parameters_general u
		JOIN data_chapter c ON u.chapter = c._id
		JOIN data_rarity r ON u.rarity = r._id
		JOIN data_type t ON u.type = t._id
		WHERE u._id = ?
	`, id).Scan(
		&u.ID,
		&u.NumID,
		&u.Chapter,
		&u.Rarity,
		&u.Type,
		&u.Transform,
		&u.LF,
		&u.Zenkai,
		&u.TagSwitch,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit: %w", err)
	}

	baseImage := unit.Images{
		BChaCut: os.Getenv("BCHACUT_URL") + strconv.Itoa(u.NumID) + ".webp",
		BChaIco: os.Getenv("BCHAICO_URL") + strconv.Itoa(u.NumID) + ".webp",
	}
	u.Images = append(u.Images, baseImage)

	if u.Transform || u.TagSwitch {
		image2 := unit.Images{
			BChaCut: os.Getenv("BCHACUT_URL") + strconv.Itoa(u.NumID) + "2.webp",
			BChaIco: os.Getenv("BCHAICO_URL") + strconv.Itoa(u.NumID) + "2.webp",
		}
		u.Images = append(u.Images, image2)
	}

	if u.Transform && u.TagSwitch {
		image3 := unit.Images{
			BChaCut: os.Getenv("BCHACUT_URL") + strconv.Itoa(u.NumID) + "3.webp",
			BChaIco: os.Getenv("BCHAICO_URL") + strconv.Itoa(u.NumID) + "3.webp",
		}
		u.Images = append(u.Images, image3)
	}

	// Add tags
	tags, err := r.GetUnitTags(u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit tags: %w", err)
	}
	u.Tags = make([]int, len(tags))

	for i, t := range tags {
		u.Tags[i] = t.ID
	}

	// Add affinity
	affinity, err := r.GetUnitAffinity(u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit affinity: %w", err)
	}
	u.Affinity = make([]int, len(affinity))

	for i, a := range affinity {
		u.Affinity[i] = a.ID
	}

	return u, nil
}

func (r *SQLUnitRepository) GetUnitNames(unitID string) ([]unit.UnitParametersName, error) {
	rows, err := r.db.Query("SELECT name_es, name_en, name_fr, name_jp FROM unit_parameters_name WHERE _unit_id = ?", unitID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit names: %w", err)
	}
	defer rows.Close()

	var names []unit.UnitParametersName
	for rows.Next() {
		var name unit.UnitParametersName
		err := rows.Scan(&name.NameES, &name.NameEN, &name.NameFR, &name.NameJP)

		if err != nil {
			return nil, fmt.Errorf("failed to scan unit name: %w", err)
		}
		names = append(names, name)
	}
	return names, nil
}

func (r *SQLUnitRepository) GetUnitTags(unitID string) ([]unit.UnitParametersTag, error) {
	rows, err := r.db.Query(`SELECT t._id, t.data_es 
		FROM unit_parameters_tag u 
		JOIN data_tag t ON u.tag = t._id
		WHERE u._unit_id = ?`, unitID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit tags: %w", err)
	}
	defer rows.Close()

	var tags []unit.UnitParametersTag
	for rows.Next() {
		var tag unit.UnitParametersTag
		err := rows.Scan(
			&tag.ID,
			&tag.DataEs,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan unit tag: %w", err)
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *SQLUnitRepository) GetUnitAffinity(unitID string) ([]unit.UnitParametersAffinity, error) {
	rows, err := r.db.Query(`
        SELECT d._id, d.data_es
        FROM unit_parameters_affinity upa
        JOIN data_affinity d ON upa.affinity = d._id
        WHERE upa._unit_id = ?;
    `,
		unitID)

	if err != nil {
		return nil, fmt.Errorf("failed to get unit affinity: %w", err)
	}
	defer rows.Close()

	var affinities []unit.UnitParametersAffinity
	for rows.Next() {
		var a unit.UnitParametersAffinity
		if err := rows.Scan(
			&a.ID,
			&a.DataEs,
		); err != nil {
			return nil, fmt.Errorf("failed to scan unit affinity: %w", err)
		}
		affinities = append(affinities, a)
	}

	return affinities, nil
}

func (r *SQLUnitRepository) GetUnitTraits(unitID string) ([]unit.UnitParametersTraits, error) {
	rows, err := r.db.Query("SELECT trait FROM unit_parameters_traits WHERE _unit_id = ?", unitID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit traits: %w", err)
	}
	defer rows.Close()

	var traits []unit.UnitParametersTraits
	for rows.Next() {
		var trait unit.UnitParametersTraits
		err := rows.Scan(&trait.Trait)
		if err != nil {
			return nil, fmt.Errorf("failed to scan unit trait: %w", err)
		}
		traits = append(traits, trait)
	}
	return traits, nil
}

func (r *SQLUnitRepository) GetUnitHeldCards(unitID string) (*unit.UnitParametersHeldCards, error) {
	heldCards := &unit.UnitParametersHeldCards{}
	err := r.db.QueryRow(`
		SELECT
			card1, card2 
		FROM unit_parameters_held_cards
		WHERE _unit_id = ?
	`, unitID).Scan(
		&heldCards.Card1,
		&heldCards.Card2,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit held cards: %w", err)
	}
	return heldCards, nil
}

func (r *SQLUnitRepository) GetUnitStatsMin(unitID string) (*unit.UnitParametersStatsMin, error) {
	stats := &unit.UnitParametersStatsMin{}
	err := r.db.QueryRow(`
		SELECT
			_id, _unit_id, health, base_strike_attack, base_blast_attack,
			base_strike_defense, base_blast_defense, critic, ki_recovery
		FROM unit_parameters_stats_min
		WHERE _unit_id = ?
	`, unitID).Scan(
		&stats.ID,
		&stats.UnitID,
		&stats.Health,
		&stats.BaseStrikeAttack,
		&stats.BaseBlastAttack,
		&stats.BaseStrikeDefense,
		&stats.BaseBlastDefense,
		&stats.Critic,
		&stats.KiRecovery,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit stats min: %w", err)
	}
	return stats, nil
}

func (r *SQLUnitRepository) GetUnitStatsMax(unitID string) (*unit.UnitParametersStatsMax, error) {
	stats := &unit.UnitParametersStatsMax{}
	err := r.db.QueryRow(`
		SELECT
			_id, _unit_id, health, base_strike_attack, base_blast_attack,
			base_strike_defense, base_blast_defense, critic, ki_recovery
		FROM unit_parameters_stats_max
		WHERE _unit_id = ?
	`, unitID).Scan(
		&stats.ID,
		&stats.UnitID,
		&stats.Health,
		&stats.BaseStrikeAttack,
		&stats.BaseBlastAttack,
		&stats.BaseStrikeDefense,
		&stats.BaseBlastDefense,
		&stats.Critic,
		&stats.KiRecovery,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get unit stats max: %w", err)
	}
	return stats, nil
}
