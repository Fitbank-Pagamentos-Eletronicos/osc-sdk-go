package domains

import "encoding/json"

type Proposal struct {
	Mother             string
	Gender             Gender
	Natianality        Nationality
	HomeTownState      HomeTownState
	HomeTownCity       string
	Education          EducationLevel
	RelationshipStatus RelationshipStatus
	Identity           IdentityType
	Adrress            Address
	Vehicle            Vehicle
	ConsumerUnit       ConsumerUnit
	Bussiness          Business
	Bank               Bank
	Reference          Reference
	Products           Products
}

func (a *Proposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
