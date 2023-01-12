package domains

import "encoding/json"

type AcceptedCheckSumBody struct {
	AcceptedCheckSum string `json:"acceptedCheckSum"`
}

func (a *AcceptedCheckSumBody) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
