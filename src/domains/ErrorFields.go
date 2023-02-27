package domains

import json2 "encoding/json"

type ErrorFields struct {
	Code    string
	Message string
	Errors  ErrorField
}

func (a *ErrorFields) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
