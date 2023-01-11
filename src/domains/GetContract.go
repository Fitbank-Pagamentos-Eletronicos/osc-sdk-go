package domains

import "encoding/json"

type GetContract struct {
	CustomerServiceNumber string `json:"customerServiceNumber"`
	Contracts             []ContractBody `json:"contracts"`
}

func (a *GetContract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
