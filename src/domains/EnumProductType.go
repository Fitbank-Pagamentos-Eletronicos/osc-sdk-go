package domains

type ProductType struct{
	productType string
}

const (
	LOAN ProductType = ProductType{"LOAN"}
	CARD  ProductType = ProductType{"CARD"}
	REFINANCING_AUTO ProductType = ProductType{"REFINANCING_AUTO"}
	REFINANCING_HOME ProductType = ProductType{"REFINANCING_HOME"}
)