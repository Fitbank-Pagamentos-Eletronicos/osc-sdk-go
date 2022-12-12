package domains

import "encoding/json"

type ProposalBankAccount struct {
	Mother             string
	Gender             Gender
	Nationality        Nationality
	HomeTownState      HomeTownState
	RelationshipStatus RelationshipStatus
	Identity           IdentityType
	Address             Address
	Business           Business
	Products           ProductBankAccount
}

func (a *ProposalBankAccount) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
