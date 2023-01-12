package domains

import "encoding/json"

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
	json, _ := json.Marshal(a)
	return string(json)
}
