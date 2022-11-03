package domains

type PipelineWithProposal struct {
	id string
	status string
	cpf string
	name string
	dateCreated string
	lastUpdated string
	proposals Proposals
}