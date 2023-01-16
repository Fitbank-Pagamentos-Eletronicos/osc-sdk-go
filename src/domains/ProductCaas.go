package domains

import json2 "encoding/json"

type ProductCaas struct {
	Type         ProductType `json:"type"`
	ReleaseDate  string      `json:"releaseDate"`
	DueDate      string      `json:"dueDate"`
	Installments int         `json:"installments"`
	MonthlyTax   float64     `json:"monthlyTax"`
	CadValue     float64     `json:"cadValue"`
	Value        float64     `json:"value"`
}

func (a *ProductCaas) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
