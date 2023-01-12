package domains

import "encoding/json"

type ConsumerUnit struct {
	Number string `json:"number"`
}

func (a *ConsumerUnit) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
