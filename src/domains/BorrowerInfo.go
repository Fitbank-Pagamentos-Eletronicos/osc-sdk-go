package domains

import json2 "encoding/json"

type BorrowerInfo struct {
	IsRegistered bool `json:"isRegistered"`
	IsBlocked    bool `json:"isBlocked"`
}

func (a *BorrowerInfo) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
