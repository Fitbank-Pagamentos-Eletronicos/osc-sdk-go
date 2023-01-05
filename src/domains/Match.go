package domains

import "encoding/json"

type Match struct {
	//Loan, Auto and Home
	ProductId      int    `json:"productId"`
	Name           string `json:"name"`
	Logo           string `json:"logo"`
	MinValue       int    `json:"minValue"`
	MaxValue       int    `json:"maxValue"`
	MinInstallment int    `json:"minInstallment"`
	MaxInstallment int
	MonthlyTax     float64 `json:"monthlyTax"`

	//Card
	Annuity float64 `json:"annuity"`
	Network NetWork `json:"network"`
}

func (a *Match) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
