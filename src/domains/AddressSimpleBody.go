package domains

import json2 "encoding/json"

type AddressSimpleBody struct {
	ZipCode    string        `json:"zipCode"`
	Address    string        `json:"address"`
	Number     string        `json:"number"`
	Complement string        `json:"complement"`
	District   string        `json:"district"`
	State      HomeTownState `json:"state"`
	City       string        `json:"city"`
}

func (a *AddressSimpleBody) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
