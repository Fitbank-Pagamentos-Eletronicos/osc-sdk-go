package domains

import "encoding/json"

type ErrorFields struct {
	Code    string
	Message string
	Errors  ErrorField
}

func (a *ErrorFields) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
