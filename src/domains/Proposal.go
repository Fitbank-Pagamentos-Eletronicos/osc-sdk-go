package domains

import "encoding/json"

type Proposal struct {
	Mother             string `json:"mother"`
	Gender             Gender `json:"gender"`
	Nationality        Nationality `json:"nationality"`
	HomeTownState      HomeTownState `json:"hometownState"`
	HomeTownCity       string `json:"hometown"`
	Education          EducationLevel `json:"education"`
	RelationshipStatus RelationshipStatus `json:"relationshipStatus"`
	PhoneLandLine      string `json:"phoneLandline"`
	Identity           Identity `json:"identity"`
	Address            Address  `json:"address"`
	Vehicle            Vehicle `json:"vehicle"`
	ConsumerUnit       ConsumerUnit `json:"consumerUnit"`
	Business           Business `json:"business"`
	Bank               Bank `json:"bank"`
	Reference          Reference
	Products           Products
}

func (a *Proposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
