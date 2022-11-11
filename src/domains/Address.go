package domains


import (
	"encoding/json"
)
type Address struct {
	ZipCode string
	Address string
	Number string
	Complement string
	District string
	State HomeTownState
	City string
	HomeType_ HomeType
	HomeSince_ HomeSince
}

func (a *Address) ToJson() string{
	json, _ := json.Marshal(a)
	return string(json)
}

