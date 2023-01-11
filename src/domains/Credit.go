package domains

import "encoding/json"

type Credit struct {
	CustomerServiceNumber string           `json:"customerServiceNumber"`
	Type                  ProductType      `json:"tipo"`
	Product               string           `json:"product"`
	ProductId             int              `json:"productId"`
	HasDocuments          bool             `json:"hasDocuments"`
	HasContracts          bool             `json:"hasContracts"`
	Logo                  string           `json:"logo"`
	LastStatus            CreditStatus     `json:"lastStatus"`
	PendentDocuments      PendentDocuments `json:"pendentDocuments"`
	DateCreated           string           `json:"dateCreated"`
	LastUpdated           string           `json:"lastUpdated"`
}

func (a *Credit) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
