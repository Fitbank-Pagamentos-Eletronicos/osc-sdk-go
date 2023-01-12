package domains

import json2 "encoding/json"

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
	json, _ := json2.Marshal(a)
	return string(json)
}
