package domains

import "encoding/json"

type Loan struct {
	CustomerServiceNumber string
	Tipo                  ProductType
	Product               string
	ProductId             int
	HasDocuments          bool
	HasContracts          bool
	Logo                  string
	LastStatus            CreditStatus
	PendentDocuments      PendentDocuments
	DateCreated           string
	LastUpdated           string
	Value                 float64
	Installments          int
	MonthlyTax            float64
	InstallmentsValue     float64
	IofValeu              float64
	GrossValeu            float64
	FirstPaymentDate      string
	Cet                   float64
	ReleaseDate           string
}

func (a *Loan) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
