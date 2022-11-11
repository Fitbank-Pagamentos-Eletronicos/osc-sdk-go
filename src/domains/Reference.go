package domains

import "encoding/json"

type Reference struct {
	Name  string
	Phone string
}

func (a *Reference) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
