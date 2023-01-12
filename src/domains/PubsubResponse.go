package domains

import json2 "encoding/json"

type PubsubResponse struct {
	Topic_id        string `json:"topic_id"`
	Subscription_id string `json:"subscription_id"`
	Project_id      string `json:"project_id"`
	Service_account string `json:"service_account"`
}

func (a *PubsubResponse) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
