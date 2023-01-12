package domains

import json2 "encoding/json"

type BusinessSimpleBody struct {
	Occupation Occupation `json:"occupation"`
	Income     string     `json:"income"`
}

func (a *BusinessSimpleBody) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
