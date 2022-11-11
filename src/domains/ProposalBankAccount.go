package domains

import "encoding/json"

type ProposalBankAccount struct {
	Mother             string
	Gender             Gender
	Natianality        Nationality
	HomeTownState      HomeTownState
	RelationshipStatus RelationshipStatus
	Identity           IdentityType
	Adrres             Address
	Business           Business
	Products           ProductBankAccount
}

func (a *ProposalBankAccount) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
