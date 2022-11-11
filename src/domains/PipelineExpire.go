package domains

import "encoding/json"

type PipelineExpire struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
}

func (a *PipelineExpire) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
