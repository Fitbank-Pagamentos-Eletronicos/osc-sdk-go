package domains

import json2 "encoding/json"

type ProductLoan struct {
	Type         ProductType `json:"type"`
	Value        int         `json:"value"`
	Installments int         `json:"installments"`
}

func (a *ProductLoan) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
