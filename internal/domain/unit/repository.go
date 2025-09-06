package unit

type Repository interface {
	GetUnitByID(id string) (*Unit, error)
	// CreateUnit(unit *Unit) error
	// UpdateUnit(unit *Unit) error
	// DeleteUnit(id string) error
	// GetAllUnitSummaries() ([]*UnitSummary, error)
	GetAllUnitSummariesCached() ([]*UnitSummary, error) 
	GetUnitSummaryByID(id string) (*UnitSummary, error)
	GetUnitNames(unitID string) ([]UnitParametersName, error)
	GetUnitTags(unitID string) ([]UnitParametersTag, error)
	GetUnitAffinity(unitID string) ([]UnitParametersAffinity, error)
	GetUnitTraits(unitID string) ([]UnitParametersTraits, error)
	GetUnitHeldCards(unitID string) (*UnitParametersHeldCards, error)
	GetUnitStatsMin(unitID string) (*UnitParametersStatsMin, error)
	GetUnitStatsMax(unitID string) (*UnitParametersStatsMax, error)
}
