package domains

import "encoding/json"

type Product struct {
	ProductId int
	Name      string
	Logo      string
}

func (a *Product) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
