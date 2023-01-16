package domains

import json2 "encoding/json"

type ProductCard struct {
	Type    ProductType `json:"type"`
	Network NetWork     `json:"network"`
	Payday  string      `json:"payday"`
}

func (a *ProductCard) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
