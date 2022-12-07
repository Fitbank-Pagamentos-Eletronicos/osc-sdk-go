package domains

import "encoding/json"

type Bank struct {
	Bank   Banks
	Tipo    AccountType
	Agency  string
	Account string
}

func (a *Bank) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
