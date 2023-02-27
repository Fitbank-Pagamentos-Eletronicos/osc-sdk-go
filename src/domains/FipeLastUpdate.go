package domains

import json2 "encoding/json"

type FipeLastUpdate struct {
	LastUpdate string `json:"lastUpdate"`
}

func (a *FipeLastUpdate) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
