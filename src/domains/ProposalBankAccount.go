package domains

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
