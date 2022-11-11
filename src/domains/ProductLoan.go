package domains

import "encoding/json"

type ProductLoan struct {
	Tipo         ProductType
	Value        int
	Installments int
}

func (a *ProductLoan) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
