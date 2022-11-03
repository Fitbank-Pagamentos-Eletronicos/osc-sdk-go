package domains

type Auto struct {
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
	value float64
	installments int16
	monthlyTax float64
	installmentsValue float64
	iofValeu float64
	grossValeu	float64
	firstPaymentDate string
	cet float64
	releaseDate string
}