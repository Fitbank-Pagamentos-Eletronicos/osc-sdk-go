package domains

import json2 "encoding/json"

type Products struct {
	//Loan
	Type         ProductType `json:"type"`
	Value        float64     `json:"value"`
	Installments int         `json:"installments"`
	//Card

	Network NetWork `json:"network"`
	Payday  string  `json:"payday"`

	// Auto
	VehicleBrand     string  `json:"vehicleBrand"`
	VehicleModel     string  `json:"vehicleModel"`
	VehicleModelYear string  `json:"vehicleModelYear"`
	CodeFipe         string  `json:"codeFipe"`
	VehicleFipeValue float64 `json:"vehicleFipeValue"`

	//Home
	RealEstateType     RealEstateType `json:"realEstateType"`
	RealEstateValue    float64        `json:"realEstateValue"`
	OutstandingBalance float64        `json:"outstandingBalance"`
}

func (a *Products) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
