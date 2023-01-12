package domains

import json2 "encoding/json"

type GetContract struct {
	CustomerServiceNumber string         `json:"customerServiceNumber"`
	Contracts             []ContractBody `json:"contracts"`
}

func (a *GetContract) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
