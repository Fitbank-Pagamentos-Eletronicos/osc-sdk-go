package domains

import "encoding/json"

type Contract struct {
	AceptedCheckSum string
	LogData_        LogData
}

func (a *Contract) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
