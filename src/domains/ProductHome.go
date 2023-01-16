package domains

import json2 "encoding/json"

type ProductHome struct {
	Type               ProductType    `json:"type"`
	Value              float64        `json:"value"`
	Installments       int            `json:"installments"`
	RealEstateType     RealEstateType `json:"realEstateType"`
	RealEstateValue    float64        `json:"realEstateValue"`
	OutstandingBalance float64        `json:"outstandingBalance"`
}

func (a *ProductHome) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
