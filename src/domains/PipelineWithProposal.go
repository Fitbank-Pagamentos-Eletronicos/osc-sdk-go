package domains

type PipelineWithProposal struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
	Proposals_  Proposals
}
