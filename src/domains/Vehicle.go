package domains

import "encoding/json"

type Vehicle struct {
	LicensePlate string
}

func (a *Vehicle) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
