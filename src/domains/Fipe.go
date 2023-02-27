package domains

import json2 "encoding/json"

type Fipe struct {
	//description: "Veiculos da tabela fipe"
	LastUpdated string        `json:"lastUpdated"`
	Vehicle     []VehicleBody `json:"vehicle"`
}

func (a *Fipe) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
