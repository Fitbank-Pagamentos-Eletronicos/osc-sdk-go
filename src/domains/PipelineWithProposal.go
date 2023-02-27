package domains

import json2 "encoding/json"

type PipelineWithProposal struct {
	Id          string   `json:"id"`
	Status      string   `json:"status"`
	Cpf         string   `json:"cpf"`
	Name        string   `json:"name"`
	DateCreated string   `json:"dateCreated"`
	LastUpdated string   `json:"lastUpdated"`
	Proposals   Proposal `json:"proposals"`
}

func (a *PipelineWithProposal) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
