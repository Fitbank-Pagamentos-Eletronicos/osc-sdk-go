package domains

import (
	json2 "encoding/json"
	"time"
)

type ResponseConsultBankValues struct {
	Document         string    `json:"document"`
	BankAccount      string    `json:"bankAccount"`
	BankAccountDigit string    `json:"bankAccountDigit"`
	BankBranch       string    `json:"bankBranch"`
	Bank             Banks     `json:"bank"`
	StartDate        time.Time `json:"StartDate"`
	EndDate          time.Time `json:"endDate"`
	Type             string    `json:"type"`
}

func (a *ResponseConsultBankValues) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
