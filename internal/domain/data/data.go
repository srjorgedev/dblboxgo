package data

type Data struct {
	ID     int    `json:"id" db:"_id"`
	DataEs string `json:"data_es" db:"data_es"`
	DataEn string `json:"data_en" db:"data_en"`
	DataFr string `json:"data_fr" db:"data_fr"`
	DataJp string `json:"data_jp" db:"data_jp"`
}

type AllData struct {
	DataTag       []*Data `json:"data_tag"`
	DataChapter   []*Data `json:"data_chapter"`
	DataRaritie   []*Data `json:"data_raritie"`
	DataType      []*Data `json:"data_type"`
	DataAffinitie []*Data `json:"data_affinity"`
}
