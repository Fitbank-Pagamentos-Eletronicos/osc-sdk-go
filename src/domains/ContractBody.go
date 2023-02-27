package domains

import json2 "encoding/json"

type ContractBody struct {
	CheckSum string `json:"checkSum"`
	Contract string `json:"contract"`
}

func (a *ContractBody) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
