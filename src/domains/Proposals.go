package domains

import "encoding/json"

type Proposals struct {
	Loan Loan
	Card Card
	Auto Auto
	Home Home
}

func (a *Proposals) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
