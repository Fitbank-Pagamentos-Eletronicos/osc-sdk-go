package domains

import json2 "encoding/json"

type ResponseTransfer struct {
	ID               string      `json:"id"`
	CcbNumber        string      `json:"ccbNumber"`
	Value            string      `json:"value"`
	BankAccount      string      `json:"bankAccount"`
	BankAccountDigit string      `json:"bankAccountDigit"`
	BankBranch       string      `json:"bankBranch"`
	Bank             Banks       `json:"bank"`
	AccountType      AccountType `json:"accountType"`
	Document         string      `json:"document"`
	Status           string      `json:"status"`
	CreditID         string      `json:"creditId"`
}

func (a *ResponseTransfer) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
