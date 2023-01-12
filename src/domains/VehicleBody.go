package domains

import "encoding/json"

type VehicleBody struct {
	VehicleBrand     string `json:"vehicleBrand"`
	VehicleModel     string `json:"vehicleModel"`
	CodeFipe         string `json:"codeFipe"`
	VehicleFipeValue string `json:"vehicleFipeValue"`
	VehicleType      string `json:"vehicleType"`
	VehicleYear      string `json:"vehicleYear"`
}

func (a *VehicleBody) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
