package domains

import json2 "encoding/json"

type ProductBankAccount struct {
	Type ProductType `json:"type"`
}

func (a *ProductBankAccount) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
