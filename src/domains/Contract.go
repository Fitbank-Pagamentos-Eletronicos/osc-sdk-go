package domains

import json2 "encoding/json"

type Contract struct {
	AcceptedCheckSum []string `json:"acceptedCheckSum"`
	LogData          LogData  `json:"logData"`
}

func (a *Contract) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
