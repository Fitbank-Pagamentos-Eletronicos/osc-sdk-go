package domains

import "encoding/json"

type PipelineProposal struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
}

func (a *PipelineProposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
