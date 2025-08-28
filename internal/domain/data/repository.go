package data

type Repository interface {
	ReadAllTags() ([]*Data, error)
	ReadAllChapters() ([]*Data, error)
	ReadAllRarities() ([]*Data, error)
	ReadAllTypes() ([]*Data, error)
	ReadAllAffinities() ([]*Data, error)
	GetCachedTags() ([]*Data, error)
	GetCachedChapters() ([]*Data, error)
	GetCachedRarities() ([]*Data, error)
	GetCachedTypes() ([]*Data, error)
	GetCachedAffinities() ([]*Data, error)
}
