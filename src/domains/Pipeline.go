package domains

import "encoding/json"

type Pipeline struct {
	Id          string
	Status      string
	Cpf         string
	Name        string
	DateCreated string
	LastUpdated string
}

func (a *Pipeline) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
