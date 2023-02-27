package domains

import json2 "encoding/json"

type SignupMatchResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Cpf         string `json:"cpf"`
	DataCreated string `json:"dataCreated"`
	LastUpdate  string `json:"lastUpdate"`
}

func (a *SignupMatchResponse) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
