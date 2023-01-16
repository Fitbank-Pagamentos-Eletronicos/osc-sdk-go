package domains

import json2 "encoding/json"

type SingContract struct {
	AcceptedCheckSum string `json:"acceptedCheckSum"`
}

func (a *SingContract) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
