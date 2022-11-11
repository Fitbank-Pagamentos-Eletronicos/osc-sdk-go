package domains

import "encoding/json"

type MatchAuto struct {
	ProductId      int
	Name           string
	Logo           string
	MinValue       int
	MaxValue       int
	MinInstallment int
	MaxInstallment int
	MonthlyTax     float64
}

func (a *MatchAuto) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
