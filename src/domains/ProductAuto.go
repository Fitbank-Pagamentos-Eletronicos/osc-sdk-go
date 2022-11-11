package domains

import "encoding/json"

type ProductAuto struct {
	Tipo             ProductType
	Value            float64
	VehicleBrand     string
	VehicleModel     string
	Installments     int
	CodeFipe         string
	VehicleFipeValue float64
}

func (a *ProductAuto) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
