package domains

type Credit struct {
	CustomerServiceNumber string
	Tipo ProductType
	Product string
	ProductId int
	HasDocuments bool
	HasContracts bool
	Logo string
	LastStatus CreditStatus
	PendentDocuments PendentDocuments
	DateCreated string
	LastUpdated string
}