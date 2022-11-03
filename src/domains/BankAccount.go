package domains

type BankAccount struct {
	customerServiceNumber string
	product string
	productId int16
	hasDocuments bool
	hasContracts bool
	lastStatus CreditStatus
	dateCreated string
	lastUpdated string
}