package domains

import "encoding/json"

type SingContract struct {
	AceptedCheckSum string `json:"aceptedCheckSum"`

}

func (a *SingContract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
