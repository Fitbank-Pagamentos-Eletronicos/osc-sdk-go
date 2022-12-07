package domains


import (
	"encoding/json"
)
type Address struct {
	ZipCode string `json:"zipCode"`
	Address string `json:"address"`
	Number string `json:"number"`
	Complement string `json:"complement"`
	District string `json:"district"`
	State HomeTownState
	City string     `json:"city"`
	HomeType HomeType
	HomeSince HomeSince
}

func (a *Address) ToJson() string{
	json, _ := json.Marshal(a)
	return string(json)
}

