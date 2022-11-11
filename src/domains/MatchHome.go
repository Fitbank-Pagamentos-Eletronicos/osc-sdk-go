package domains

import "encoding/json"

type MatchHome struct {
	ProductId      int
	Name           string
	Logo           string
	MinValue       int
	MaxValue       int
	MinInstallment int
	MaxInstallment int
	MonthlyTax     float64
}

func (a *MatchHome) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
