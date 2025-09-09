package equipment

type Equipment struct {
	Details EquipmentDetails `db:"" json:"details"`
	Slots   EquipmentSlots   `db:"" json:"slots"`
}

type EquipmentSummary struct {
	Details EquipmentDetails `db:"" json:"details"`
}

type EquipmentDetails struct {
	ID              int             `db:"_id" json:"_id"`
	Name            EquipmentName   `db:"" json:"name"`
	EquipmentRarity int             `db:"equipment_rarity" json:"equipment_rarity"`
	IsAwakened      bool            `db:"awaken" json:"is_awakened"`
	AwakenFrom      int             `db:"awaken_from" json:"awaken_from"`
	Images          EquipmentImages `db:"" json:"images"`

	Traits []EquipmentTraits `db:"" json:"traits"`
}

type EquipmentSlots struct {
	Slot1 []string `db:"slot_1" json:"slot_1"`
	Slot2 []string `db:"slot_2" json:"slot_2"`
	Slot3 []string `db:"slot_3" json:"slot_3"`
	Slot4 []string `db:"slot_4" json:"slot_4"`
}

type EquipmentTraits struct {
	ID     int    `db:"_id" json:"_id"`
	Trait1 string `db:"trait_1" json:"trait_1"`
	Trait2 string `db:"trait_2" json:"trait_2"`
	Trait3 string `db:"trait_3" json:"trait_3"`
}

type EquipmentImages struct {
	RarityImage string `json:"rarity_image"`
	IconImage   string `json:"icon_image"`
}

type EquipmentName struct {
	NameEN string `db:"name_en" json:"name_en"`
	NameES string `db:"name_es" json:"name_es"`
	NameFR string `db:"name_fr" json:"name_fr"`
	NameJP string `db:"name_jp" json:"name_jp"`
}
