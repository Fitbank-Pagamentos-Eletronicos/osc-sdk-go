package domains

import json2 "encoding/json"

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
	json, _ := json2.Marshal(a)
	return string(json)
}
