package domains

type PipelineWithBankAccountProposal struct {
	id string
	status string
	cpf string
	name string
	dateCreated string
	lastUpdated string
	proposals BankAccount
}