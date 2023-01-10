package domains

import "encoding/json"

type PipelineWithProposal struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
	Proposals   Proposal
}

func (a *PipelineWithProposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
