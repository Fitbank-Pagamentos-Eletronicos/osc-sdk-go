package domains

import (
	json2 "encoding/json"
)

// Struct to store the OAuth token
type Auth struct {
	ClientId     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
}

func (a *Auth) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
