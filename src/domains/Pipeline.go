package domains

import "encoding/json"

type Pipeline struct {
	Id          string
	Status      PipelineStatus
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
	Matches     Match
	Proposals   []Proposal
}

func (a *Pipeline) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
