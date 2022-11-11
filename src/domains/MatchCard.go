package domains

import "encoding/json"

type MatchCard struct {
	ProductId int
	Name      string
	Logo      string
	Annuity   float64
	Network   NetWork
}

func (a *MatchCard) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
