package domains

import json2 "encoding/json"

type Vehicle struct {
	LicensePlate string `json:"licensePlate"`
}

func (a *Vehicle) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
