package domains

import "encoding/json"

type ProductHome struct {
	Tipo               ProductType
	Value              float64
	Installments       int
	RealEstateType_    RealEstateType
	RealEstateValue    float64
	OutstandingBalance float64
}

func (a *ProductHome) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
