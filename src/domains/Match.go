package domains

import "encoding/json"

type Match struct {
	MatchLoan MatchLoan
	MatchCard MatchCard
	MatchAuto MatchAuto
	MatchHome MatchHome
}

func (a *Match) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
