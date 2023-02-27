package domains

import json2 "encoding/json"

type Identity struct {
	Type        IdentityType       `json:"type"`
	Number      string             `json:"number"`
	Issuer      IdentityTypeIssuer `json:"issuer"`
	State       HomeTownState      `json:"state"`
	IssuingDate string             `json:"issuingDate"`
}

func (a *Identity) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
