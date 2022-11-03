package domains

type Card struct {
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
	international string
	annuity string
	network NetWork
	prepaid bool
	digitalAccount bool
}