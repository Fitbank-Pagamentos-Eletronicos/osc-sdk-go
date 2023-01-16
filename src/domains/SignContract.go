package domains

import json2 "encoding/json"

type SignContract struct {
	AcceptedCheckSum []AcceptedCheckSumBody `json:"acceptedCheckSum"`
}

func (a *SignContract) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
