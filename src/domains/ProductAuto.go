package domains

import json2 "encoding/json"

type ProductAuto struct {
	Type             ProductType `json:"type"`
	Value            float64     `json:"value"`
	VehicleBrand     string      `json:"vehicleBrand"`
	VehicleModel     string      `json:"vehicleModel"`
	Installments     int         `json:"installments"`
	VehicleModelYear string      `json:"vehicleModelYear"`
	CodeFipe         string      `json:"codeFipe"`
	VehicleFipeValue float64     `json:"vehicleFipeValue"`
}

func (a *ProductAuto) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
