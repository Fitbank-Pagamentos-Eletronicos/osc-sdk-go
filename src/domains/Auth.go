package domains

import (
	"encoding/json"
)

// Struct to store the OAuth token
type Auth struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Scopes        []string `json:"scopes"`
}

func (a *Auth) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
