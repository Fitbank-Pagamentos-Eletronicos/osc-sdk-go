package domains

import json2 "encoding/json"

type SimpleSignup struct {
	Cpf            string  `json:"cpf"`
	Name           string  `json:"name"`
	Birthday       string  `json:"birthday"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	ZipCode        string  `json:"zipCode"`
	HasCreditCart  bool    `json:"hasCreditCart"`
	HasRestriction bool    `json:"hasRestriction"`
	HasOwnHouse    bool    `json:"hasOwnHouse"`
	HasVehicle     bool    `json:"hasVehicle"`
	HasAndroid     bool    `json:"hasAndroid"`
	LogData        LogData `json:"logData"`
}

func (a *SimpleSignup) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
