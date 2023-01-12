package domains

import json2 "encoding/json"

type Reference struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (a *Reference) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
