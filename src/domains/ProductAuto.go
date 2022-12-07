package domains

import "encoding/json"

type ProductAuto struct {
	Tipo             ProductType
	Value            float64 `json:"value"`
	VehicleBrand     string     `json:"vehicleBrand"`
	VehicleModel     string    `json:"vehicleModel"`
    Installments     int      `json:"installments"`
	VehicleModelYear string   `json:"vehicleModelYear"`
	CodeFipe         string  `json:"codeFipe"`
	VehicleFipeValue float64 `json:"vehicleFipeValue"`
}

func (a *ProductAuto) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
