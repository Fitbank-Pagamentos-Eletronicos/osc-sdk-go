package domains

import json2 "encoding/json"

type ConsumerUnit struct {
	Number string `json:"number"`
}

func (a *ConsumerUnit) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
