package domains

import json2 "encoding/json"

type MatchCard struct {
	ProductId int     `json:"productId"`
	Name      string  `json:"name"`
	Logo      string  `json:"logo"`
	Annuity   float64 `json:"annuity"`
	Network   NetWork `json:"network"`
}

func (a *MatchCard) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
