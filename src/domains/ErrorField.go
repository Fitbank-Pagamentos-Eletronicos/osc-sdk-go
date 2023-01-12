package domains

import json2 "encoding/json"

type ErrorField struct {
	Field   string
	Message string
}

func (a *ErrorField) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
