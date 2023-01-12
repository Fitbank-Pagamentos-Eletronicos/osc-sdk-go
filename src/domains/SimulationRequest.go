package domains

import (
	"encoding/json"
	"time"
)

type SimulationRequest struct {
	Type         ProductType `json:"type"`
	ReleaseDate  time.Time   `json:"releaseDate"`
	DueDate      time.Time   `json:"dueDate"`
	Installments int         `json:"installments"`
	MonthlyTax   float64     `json:"monthlyTax"`
	CadValue     float64     `json:"cadValue"`
	Value        float64     `json:"value"`
}

func (a *SimulationRequest) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
