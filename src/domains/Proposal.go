package domains

import "encoding/json"

type Proposal struct {
	Mother             string
	Gender             Gender
	Nationality        Nationality
	HomeTownState      HomeTownState
	HomeTownCity       string
	Education          EducationLevel
	RelationshipStatus RelationshipStatus
	PhoneLandLine      string `json:"phoneLandline"`
	Identity           Identity
	Address            Address
	Vehicle            Vehicle
	ConsumerUnit       ConsumerUnit
	Business          Business
	Bank               Bank
	Reference          Reference
	Products           Products
}

func (a *Proposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
