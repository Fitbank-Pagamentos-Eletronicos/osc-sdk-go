package domains

import json2 "encoding/json"

type ProposalBankAccount struct {
	Mother             string             `json:"mother"`
	Gender             Gender             `json:"gender"`
	Nationality        Nationality        `json:"national"`
	HomeTownState      HomeTownState      `json:"homeTownState"`
	RelationshipStatus RelationshipStatus `json:"relationshipStatus"`
	Identity           IdentityType       `json:"identity"`
	Address            Address            `json:"address"`
	Business           Business           `json:"business"`
	Products           ProductBankAccount `json:"products"`
}

func (a *ProposalBankAccount) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
