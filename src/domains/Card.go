package domains

type Card struct {
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
	International         string
	Annuity               string
	Network               NetWork
	Prepaid               bool
	DigitalAccount        bool
}
