package domains

import (
	json2 "encoding/json"
	"time"
)

type Proposal struct {
	CustomerServiceNumber string           `json:"customerServiceNumber"`
	Type                  ProductType      `json:"type"`
	Product               string           `json:"product"`
	ProductId             int              `json:"productId"`
	HasDocuments          bool             `json:"hasDocuments"`
	HasContracts          bool             `json:"hasContracts"`
	Logo                  string           `json:"logo"`
	LastStatus            CreditStatus     `json:"lastStatus"`
	DateCreated           time.Time        `json:"dateCreated"`
	LastUpdated           time.Time        `json:"lastUpdated"`
	PendentDocuments      PendentDocuments `json:"pendentDocuments"`

	// Loan, Auto, Home
	Value             int       `json:"value"`
	Installments      int       `json:"installments"`
	MonthlyTax        float64   `json:"monthlyTax"`
	InstallmentsValue float64   `json:"installmentsValue"`
	IofValue          float64   `json:"iofValue"`
	GrossValue        float64   `json:"grossValue"`
	FirstPaymentDate  time.Time `json:"firstPaymentDate"`
	Cet               float64   `json:"cet"`
	ReleasedDate      time.Time `json:"releasedDate"`

	// Card
	International  bool    `json:"international"`
	Annuity        float64 `json:"annuity"`
	Network        string  `json:"network"`
	Prepaid        bool    `json:"prepaid"`
	DigitalAccount bool    `json:"digitalAccount"`
}

func (a *Proposal) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
