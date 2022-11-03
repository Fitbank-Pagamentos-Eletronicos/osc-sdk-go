package domains

type Credit struct {
	customerServiceNumber string
	tipo ProductType
	product string
	productId int16
	hasDocuments bool
	hasContracts bool
	logo string
	lastStatus CreditStatus
	pendentDocuments PendentDocuments
	dateCreated string
	lastUpdated string
}