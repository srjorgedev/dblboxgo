package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/srjorgedev/dblboxgo/internal/domain/data"
)

type SQLDataRepository struct {
	db        *sql.DB
	cache     *data.AllData
	cacheLock sync.RWMutex
}

func NewSQLDataRepository(db *sql.DB) *SQLDataRepository {
	repo := &SQLDataRepository{db: db}

	if _, err := repo.CacheData(); err != nil {
		fmt.Println("warning: could not warm up cache:", err)
	}

	return repo
}

func (r *SQLDataRepository) CacheData() (*data.AllData, error) {
	r.cacheLock.RLock()
	if r.cache != nil {
		defer r.cacheLock.RUnlock()
		return r.cache, nil
	}
	r.cacheLock.RUnlock()

	data, err := r.GetAllData()
	if err != nil {
		return nil, err
	}

	r.cacheLock.Lock()
	r.cache = data
	r.cacheLock.Unlock()

	return data, nil
}

func (r *SQLDataRepository) GetAllData() (*data.AllData, error) {
	tags, err := r.GetAllTags()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Cached tags: %d\n", len(tags))

	chapters, err := r.GetAllChapters()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Cached chapters: %d\n", len(chapters))

	rarities, err := r.GetAllRarities()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Cached rarities: %d\n", len(rarities))

	types, err := r.GetAllTypes()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Cached types: %d\n", len(types))

	affinities, err := r.GetAllAffinities()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Cached affinities: %d\n", len(affinities))

	allData := &data.AllData{
		DataTag:       tags,
		DataChapter:   chapters,
		DataRaritie:   rarities,
		DataType:      types,
		DataAffinitie: affinities,
	}

	return allData, nil
}

func (r *SQLDataRepository) GetAllTags() ([]*data.Data, error) {
	return r.GetDataFrom(r.db, "tag")
}

func (r *SQLDataRepository) GetAllChapters() ([]*data.Data, error) {
	return r.GetDataFrom(r.db, "chapter")
}

func (r *SQLDataRepository) GetAllRarities() ([]*data.Data, error) {
	return r.GetDataFrom(r.db, "rarity")
}

func (r *SQLDataRepository) GetAllTypes() ([]*data.Data, error) {
	return r.GetDataFrom(r.db, "type")
}

func (r *SQLDataRepository) GetAllAffinities() ([]*data.Data, error) {
	return r.GetDataFrom(r.db, "affinity")
}

func (r *SQLDataRepository) GetDataFrom(db *sql.DB, table string) ([]*data.Data, error) {
	query := fmt.Sprintf(`
	SELECT
	_id,
	COALESCE(data_es, ''), 
	COALESCE(data_en, ''), 
	COALESCE(data_fr, ''), 
	COALESCE(data_jp, '')  
	FROM data_%s`, table)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*data.Data
	for rows.Next() {
		var d data.Data
		if err := rows.Scan(&d.ID, &d.DataEs, &d.DataEn, &d.DataFr, &d.DataJp); err != nil {
			return nil, err
		}
		results = append(results, &d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *SQLDataRepository) GetCachedData() (*data.AllData, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache, nil
}

func (r *SQLDataRepository) GetCachedTags() ([]*data.Data, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache.DataTag, nil
}

func (r *SQLDataRepository) GetCachedChapters() ([]*data.Data, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache.DataChapter, nil
}

func (r *SQLDataRepository) GetCachedRarities() ([]*data.Data, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache.DataRaritie, nil
}

func (r *SQLDataRepository) GetCachedTypes() ([]*data.Data, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache.DataType, nil
}

func (r *SQLDataRepository) GetCachedAffinities() ([]*data.Data, error) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()
	if r.cache == nil {
		return nil, fmt.Errorf("cache is empty")
	}
	return r.cache.DataAffinitie, nil
}

// ReadAllAffinities implements data.Repository.
func (r *SQLDataRepository) ReadAllAffinities() ([]*data.Data, error) {
	panic("unimplemented")
}

// ReadAllChapters implements data.Repository.
func (r *SQLDataRepository) ReadAllChapters() ([]*data.Data, error) {
	panic("unimplemented")
}

// ReadAllRarities implements data.Repository.
func (r *SQLDataRepository) ReadAllRarities() ([]*data.Data, error) {
	panic("unimplemented")
}

// ReadAllTags implements data.Repository.
func (r *SQLDataRepository) ReadAllTags() ([]*data.Data, error) {
	panic("unimplemented")
}

// ReadAllTypes implements data.Repository.
func (r *SQLDataRepository) ReadAllTypes() ([]*data.Data, error) {
	panic("unimplemented")
}
