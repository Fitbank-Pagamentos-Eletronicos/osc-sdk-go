package domains

import "encoding/json"

type SignContract struct {
	AcceptedCheckSum []AcceptedCheckSumBody `json:"acceptedCheckSum"`
}

func (a *SignContract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
