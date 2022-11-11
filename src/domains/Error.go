package domains

import "encoding/json"

type Error struct {
	Code    string
	Message string
}

func (a *Error) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
