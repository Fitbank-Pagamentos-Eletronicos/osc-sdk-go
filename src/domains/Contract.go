package domains

import "encoding/json"

type Contract struct {
	AceptedCheckSum []string `json:"aceptedCheckSum"`
	LogData     LogData `json:"logData"`
}

func (a *Contract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
