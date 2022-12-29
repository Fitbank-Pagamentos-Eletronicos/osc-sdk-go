package domains

import "encoding/json"

type Bank struct {
	Bank   Banks `json:"bank"`
	Tipo    AccountType `json:"type"`
	Agency  string `json:"agency"`
	Account string `json:"account"`
}

func (a *Bank) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
