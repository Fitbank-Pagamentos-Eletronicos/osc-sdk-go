package domains

import "encoding/json"

type PipelineMatchLoan struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
	Matches     Match
}

func (a *PipelineMatchLoan) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
