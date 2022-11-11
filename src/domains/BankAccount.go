package domains

import "encoding/json"

type BankAccount struct {
	CustomerServiceNumber string
	Product               string
	ProductId             int
	HasDocuments          bool
	HasContracts          bool
	LastStatus            CreditStatus
	DateCreated           string
	LastUpdated           string
}

func (a *BankAccount) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
