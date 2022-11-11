package domains

import "encoding/json"

type MatchLoan struct {
	ProductId      int
	Name           string
	Logo           string
	MinValue       int
	MaxValue       int
	MinInstallment int
	MaxInstallment int
	MonthlyTax     float64
}

func (a *MatchLoan) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
