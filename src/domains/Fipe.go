package domains

import "encoding/json"

type Fipe struct {
	//description: "Veiculos da tabela fipe"
	LastUpdated string
	Vehicles    []string
}

func (a *Fipe) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
