package domains

import json2 "encoding/json"

type Card struct {
}

func (a *Card) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
