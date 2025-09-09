package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/srjorgedev/dblboxgo/internal/domain/equipment"
	"github.com/srjorgedev/dblboxgo/pkg"
)

type SQLEquipmentRepository struct {
	db        *sql.DB
	cache     []*equipment.EquipmentSummary
	cacheLock sync.RWMutex
}

// GetEquipmentByID implements equipment.Repository.
func (r *SQLEquipmentRepository) GetEquipmentByID(id int) (*equipment.Equipment, error) {
	panic("unimplemented")
}

// GetEquipmentSummaryByID implements equipment.Repository.
func (r *SQLEquipmentRepository) GetEquipmentSummaryByID(id int) (*equipment.EquipmentSummary, error) {
	panic("unimplemented")
}

func NewSQLEquipmentRepository(db *sql.DB) *SQLEquipmentRepository {
	repo := &SQLEquipmentRepository{db: db}

	if _, err := repo.GetAllEquipmentSummariesCached(); err != nil {
		fmt.Println("[ API ] warning: could not warm up equipment cache:", err)
	}

	return repo
}

func (r *SQLEquipmentRepository) GetAllEquipmentSummariesCached() ([]*equipment.EquipmentSummary, error) {
	r.cacheLock.RLock()
	if r.cache != nil {
		defer r.cacheLock.RUnlock()
		return r.cache, nil
	}
	r.cacheLock.RUnlock()

	summaries, err := r.GetAllEquipmentSummaries()
	if err != nil {
		return nil, err
	}
	r.cacheLock.Lock()
	r.cache = summaries
	r.cacheLock.Unlock()

	return summaries, nil
}

func (r *SQLEquipmentRepository) GetAllEquipmentSummaries() ([]*equipment.EquipmentSummary, error) {
	rows, err := r.db.Query(`SELECT
  _id,
  name_en,
  name_es,
  name_fr,
  name_jp,
  equipment_rarity,
  awaken,
  COALESCE(awaken_from, 0) AS awaken_from
FROM
  equipment_general
ORDER BY
  equipment_rarity DESC,
  _id DESC`)
	if err != nil {
		return nil, fmt.Errorf("failed to get equipment summaries: %w", err)
	}
	defer rows.Close()

	var equipments []*equipment.EquipmentSummary
	for rows.Next() {
		var eq equipment.EquipmentSummary

		if err := rows.Scan(
			&eq.Details.ID,
			&eq.Details.Name.NameEN,
			&eq.Details.Name.NameES,
			&eq.Details.Name.NameFR,
			&eq.Details.Name.NameJP,
			&eq.Details.EquipmentRarity,
			&eq.Details.IsAwakened,
			&eq.Details.AwakenFrom,
		); err != nil {
			return nil, fmt.Errorf("failed to scan equipment summary: %w", err)
		}

		eq.Details.Images = pkg.GetEquipmentImages(eq.Details.ID, eq.Details.EquipmentRarity, eq.Details.IsAwakened, eq.Details.AwakenFrom)

		equipments = append(equipments, &eq)
	}

	return equipments, nil
}
