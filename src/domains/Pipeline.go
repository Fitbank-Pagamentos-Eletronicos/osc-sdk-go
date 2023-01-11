package domains

import "encoding/json"

type Pipeline struct {
	Id          string         `json:"id"`
	Status      PipelineStatus `json:"status"`
	Cpf         string         `json:"cpf"`
	Name        string         `json:"name"`
	DateCreated string         `json:"dateCreated"`
	LastUpdated string         `json:"lastUpdated"`
	Matches     Match          `json:"matches"`
	Proposals   []Proposal     `json:"proposals"`
}

func (a *Pipeline) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
