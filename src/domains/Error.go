package domains

import json2 "encoding/json"

type Error struct {
	Code    string
	Message string
}

func (a *Error) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
