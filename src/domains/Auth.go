package domains

import (
	"encoding/json"
)

// Struct to store the OAuth token
type Auth struct {
	Client_id     string
	Client_secret string
	Scopes        []string
}

func (a *Auth) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
