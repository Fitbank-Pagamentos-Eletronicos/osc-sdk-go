package domains

import "encoding/json"

type ProductCard struct {
	Type    ProductType
	Network NetWork
	Payday  string
}

func (a *ProductCard) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
