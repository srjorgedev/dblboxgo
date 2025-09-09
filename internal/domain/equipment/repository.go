package equipment

type Repository interface {
	// GetAllEquipmentSummaries() ([]*EquipmentSummary, error)
	GetAllEquipmentSummariesCached() ([]*EquipmentSummary, error)
	GetEquipmentByID(id int) (*Equipment, error)
	GetEquipmentSummaryByID(id int) (*EquipmentSummary, error)
}