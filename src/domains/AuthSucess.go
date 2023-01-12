package domains

import json2 "encoding/json"

type AuthSucess struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    string `json:"expireAt"`
}

func (a *AuthSucess) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
