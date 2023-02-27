package domains

import json2 "encoding/json"

type AuthSucess struct {
	Access_token string `json:"access_token"`
	Expire_at    string `json:"expire_at"`
}

func (a *AuthSucess) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
