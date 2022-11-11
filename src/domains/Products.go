package domains

import "encoding/json"

type Products struct {
	ProductLoan ProductLoan
	ProductCard ProductCard
	ProductAuto ProductAuto
	ProductHome ProductHome
}

func (a *Products) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
