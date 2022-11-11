package domains

import "encoding/json"

type BorrowerInfo struct {
	IsRegistered bool
	IsBlocked    bool
}

func (a *BorrowerInfo) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
