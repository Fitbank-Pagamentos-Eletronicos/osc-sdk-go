package domains

import json2 "encoding/json"

type ProposalReq struct {
	Mother             string             `json:"mother"`
	Gender             Gender             `json:"gender"`
	Nationality        Nationality        `json:"nationality"`
	HomeTownState      HomeTownState      `json:"hometownState"`
	HomeTownCity       string             `json:"hometown"`
	Education          EducationLevel     `json:"education"`
	RelationshipStatus RelationshipStatus `json:"relationshipStatus"`
	PhoneLandLine      string             `json:"phoneLandline"`
	Identity           Identity           `json:"identity"`
	Address            Address            `json:"address"`
	Vehicle            Vehicle            `json:"vehicle"`
	ConsumerUnit       ConsumerUnit       `json:"consumerUnit"`
	Business           Business           `json:"business"`
	Bank               Bank               `json:"bank"`
	Reference          Reference
	Products           Products
}

func (a *ProposalReq) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
