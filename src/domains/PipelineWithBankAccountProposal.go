package domains

type PipelineWithBankAccountProposal struct {
	Id string
	Status string
	Cpf string
	Name string
	DateCreated string
	LastUpdated string
	Proposals BankAccount
}