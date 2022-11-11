package domains

import "encoding/json"

type ErrorField struct {
	Field   string
	Message string
}

func (a *ErrorField) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
