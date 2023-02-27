package domains

import json2 "encoding/json"

type Bank struct {
	Bank    Banks       `json:"bank"`
	Type    AccountType `json:"type"`
	Agency  string      `json:"agency"`
	Account string      `json:"account"`
}

func (a *Bank) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
