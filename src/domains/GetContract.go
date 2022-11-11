package domains

import "encoding/json"

type GetContract struct {
	CustomerServiceNumber string
	Contracts             []string
}

func (a *GetContract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
