package models

type ReturnData struct {
	Status string     `json:"status"`
	Data   []Schedule `json:"data"`
}

type Schedule struct {
	TRX_ID         int    `json:"trx_id"`
	Thbl           string `json:"thbl"`
	Volume         string `json:"volume"`
	Org_id_pemasok string `json:"org_id_pemasok"`
}

type Q struct {
	Thbl           string `json:"thbl"`
	Org_id_pemasok string `json:"org_id_pemasok"`
}
