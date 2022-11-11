package domains

import "encoding/json"

type SingContract struct {
	AceptedCheckSum string
}

func (a *SingContract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
