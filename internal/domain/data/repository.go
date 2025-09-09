package data

type Repository interface {
	GetCachedTags() ([]*Data, error)
	GetCachedChapters() ([]*Data, error)
	GetCachedRarities() ([]*Data, error)
	GetCachedTypes() ([]*Data, error)
	GetCachedAffinities() ([]*Data, error)
}
