package domains

import json2 "encoding/json"

type MatchHome struct {
	ProductId      int     `json:"productId"`
	Name           string  `json:"name"`
	Logo           string  `json:"logo"`
	MinValue       int     `json:"minValue"`
	MaxValue       int     `json:"maxValue"`
	MinInstallment int     `json:"minInstallment"`
	MaxInstallment int     `json:"maxInstallment"`
	MonthlyTax     float64 `json:"monthlyTax"`
}

func (a *MatchHome) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
