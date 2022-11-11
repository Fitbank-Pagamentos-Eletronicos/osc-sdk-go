package domains

import "encoding/json"

type PipelineWithBankAccountProposal struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
	Proposals   BankAccount
}

func (a *PipelineWithBankAccountProposal) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
