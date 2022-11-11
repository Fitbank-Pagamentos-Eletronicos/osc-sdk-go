package domains

import "encoding/json"

type AuthSucess struct {
	Acess_token string
	Expire_at   string
}

func (a *AuthSucess) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}

