package domains

import "encoding/json"

type BusinessSimpleBody struct {
	Occupation Occupation `json:"occupation"`
	Income     string     `json:"income"`
}

func (a *BusinessSimpleBody) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
