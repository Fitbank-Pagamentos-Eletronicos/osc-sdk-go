package domains

import "encoding/json"

type Identity struct {
	Tipo        IdentityType
	Number      string
	Issuer      IdentityTypeIssuer
	State       HomeTownState
	IssuingDate string
}

func (a *Identity) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
