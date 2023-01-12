package domains

import json2 "encoding/json"

type AcceptedCheckSumBody struct {
	AcceptedCheckSum string `json:"acceptedCheckSum"`
}

func (a *AcceptedCheckSumBody) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
