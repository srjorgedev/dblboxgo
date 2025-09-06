package unit

type Data struct {
	ID     int    `json:"id" db:"_id"`
	DataEs string `json:"data_es" db:"data_es"`
}

type Images struct {
	BChaCut string `json:"bchacut"`
	BChaIco string `json:"bchaico"`
}

type Unit struct {
	ID            string `json:"_id" db:"_id"`
	NumID         int    `json:"num_id" db:"num_id"`
	Chapter       Data   `json:"chapter" db:"chapter"`
	Rarity        Data   `json:"rarity" db:"rarity"`
	Type          Data   `json:"type" db:"type"`
	Transform     bool   `json:"transform" db:"transform"`
	LF            bool   `json:"lf" db:"lf"`
	Zenkai        bool   `json:"zenkai" db:"zenkai"`
	KitType       int    `json:"kit_type" db:"kit_type"`
	ZenkaiKitType int    `json:"zenkai_kit_type" db:"zenkai_kit_type"`
	TagSwitch     bool   `json:"tag_switch" db:"tag_switch"`

	Names     []UnitParametersName     `json:"names,omitempty"`
	Tags      []UnitParametersTag      `json:"tags,omitempty"`
	Affinity  []UnitParametersAffinity `json:"affinity,omitempty"`
	Traits    []UnitParametersTraits   `json:"traits,omitempty"`
	HeldCards *UnitParametersHeldCards `json:"held_cards,omitempty"`
	StatsMin  UnitParametersStatsMin   `json:"stats_min,omitempty"`
	StatsMax  UnitParametersStatsMax   `json:"stats_max,omitempty"`
}

type UnitSummary struct {
	ID        string   `json:"_id" db:"_id"`
	NumID     int      `json:"num_id" db:"num_id"`
	Chapter   int      `json:"chapter" db:"chapter"`
	Rarity    int      `json:"rarity" db:"rarity"`
	Type      int      `json:"type" db:"type"`
	Transform bool     `json:"transform" db:"transform"`
	LF        bool     `json:"lf" db:"lf"`
	Zenkai    bool     `json:"zenkai" db:"zenkai"`
	TagSwitch bool     `json:"tag_switch" db:"tag_switch"`
	Images    []Images `json:"images"`

	Tags     []int `json:"tags,omitempty"`
	Affinity []int `json:"affinity,omitempty"`
}

type UnitTierMaker struct {
	ID        string   `json:"_id" db:"_id"`
	NumID     int      `json:"num_id" db:"num_id"`
	Chapter   int      `json:"chapter" db:"chapter"`
	Rarity    int      `json:"rarity" db:"rarity"`
	Type      int      `json:"type" db:"type"`
	Transform bool     `json:"transform" db:"transform"`
	TagSwitch bool     `json:"tag_switch" db:"tag_switch"`
	Images    []Images `json:"images"`

	Tags     []int `json:"tags,omitempty"`
	Affinity []int `json:"affinity,omitempty"`
}

type UnitParametersGeneral struct {
	ID            string `json:"_id" db:"_id"`
	NumID         int    `json:"num_id" db:"num_id"`
	Chapter       Data   `json:"chapter" db:"chapter"`
	Rarity        Data   `json:"rarity" db:"rarity"`
	Type          Data   `json:"type" db:"type"`
	Transform     bool   `json:"transform" db:"transform"`
	LF            bool   `json:"lf" db:"lf"`
	Zenkai        bool   `json:"zenkai" db:"zenkai"`
	KitType       int    `json:"kit_type" db:"kit_type"`
	ZenkaiKitType int    `json:"zenkai_kit_type" db:"zenkai_kit_type"`
	TagSwitch     bool   `json:"tag_switch" db:"tag_switch"`
}

type UnitParametersName struct {
	Name string `db:"name" json:",inline"`
}

type UnitParametersAffinity struct {
	Data `db:"affinity" json:",inline"`
}

type UnitParametersTag struct {
	Data `db:"tag" json:",inline"`
}

type UnitParametersTraits struct {
	Trait int `db:"trait" json:"trait"`
}

type UnitParametersHeldCards struct {
	Card1 string `db:"card1" json:"card1"`
	Card2 string `db:"card2" json:"card2"`
}

type UnitParametersStatsMin struct {
	ID                int    `db:"_id" json:"_id"`
	UnitID            string `db:"_unit_id" json:"_unit_id"`
	Health            int    `db:"health" json:"health"`
	BaseStrikeAttack  int    `db:"base_strike_attack" json:"base_strike_attack"`
	BaseBlastAttack   int    `db:"base_blast_attack" json:"base_blast_attack"`
	BaseStrikeDefense int    `db:"base_strike_defense" json:"base_strike_defense"`
	BaseBlastDefense  int    `db:"base_blast_defense" json:"base_blast_defense"`
	Critic            int    `db:"critic" json:"critic"`
	KiRecovery        int    `db:"ki_recovery" json:"ki_recovery"`
}

type UnitParametersStatsMax struct {
	ID                int    `db:"_id" json:"_id"`
	UnitID            string `db:"_unit_id" json:"_unit_id"`
	Health            int    `db:"health" json:"health"`
	BaseStrikeAttack  int    `db:"base_strike_attack" json:"base_strike_attack"`
	BaseBlastAttack   int    `db:"base_blast_attack" json:"base_blast_attack"`
	BaseStrikeDefense int    `db:"base_strike_defense" json:"base_strike_defense"`
	BaseBlastDefense  int    `db:"base_blast_defense" json:"base_blast_defense"`
	Critic            int    `db:"critic" json:"critic"`
	KiRecovery        int    `db:"ki_recovery" json:"ki_recovery"`
}
