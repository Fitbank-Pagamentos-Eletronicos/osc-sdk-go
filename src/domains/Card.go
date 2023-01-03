package domains

import "encoding/json"

type Card struct {
}

func (a *Card) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
