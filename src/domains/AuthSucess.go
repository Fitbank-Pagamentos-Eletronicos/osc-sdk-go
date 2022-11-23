package domains

import "encoding/json"

type AuthSucess struct {
	Access_token string `json:"access_token"`
	Expire_at   string  `json:"expire_at"`
}

func (a *AuthSucess) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}

