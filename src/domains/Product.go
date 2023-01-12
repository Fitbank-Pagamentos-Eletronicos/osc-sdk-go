package domains

import json2 "encoding/json"

type Product struct {
	ProductId int
	Name      string
	Logo      string
}

func (a *Product) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
