package domains

type ProductType struct{
	productType string
}

const (
	LOAN = ProductType{"LOAN"}
	CARD = ProductType{"CARD"}
	REFINANCING_AUTO = ProductType{"REFINANCING_AUTO"}
	REFINANCING_HOME = ProductType{"REFINANCING_HOME"}
)