package domains

type PipelineMatchLoan struct {
	id string
	status string
	cpf string
	name string
	dateCreated string
	lastUpdated string
	matches Match
}