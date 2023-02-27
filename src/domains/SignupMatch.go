package domains

import json2 "encoding/json"

type SignupMatch struct {
	Cpf            string         `json:"cpf"`
	Name           string         `json:"name"`
	Birthday       string         `json:"birthday"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	ZipCode        string         `json:"zipCode"`
	Education      EducationLevel `json:"education"`
	Banks          Banks          `json:"banks"`
	Occupation     Occupation     `json:"occupation"`
	Income         int            `json:"income"`
	HasCreditCart  bool           `json:"hasCreditCart"`
	HasRestriction bool           `json:"hasRestriction"`
	HasOwnHouse    bool           `json:"hasOwnHouse"`
	HasVehicle     bool           `json:"hasVehicle"`
	HasAndroid     bool           `json:"hasAndroid"`
	Products       Products
	LogData        LogData `json:"logData"`
}

func (a *SignupMatch) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
