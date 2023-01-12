package domains

import json2 "encoding/json"

type RequestTransfer struct {
	Value            string      `json:"value"`
	BankAccount      string      `json:"bankAccount"`
	BankAccountDigit string      `json:"bankAccountDigit"`
	BankBranch       string      `json:"bankBranch"`
	Bank             Banks       `json:"bank"`
	AccountType      AccountType `json:"accountType"`
	Document         string      `json:"document"`
}

func (a *RequestTransfer) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
