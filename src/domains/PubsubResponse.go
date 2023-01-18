package domains

import json2 "encoding/json"

type PubsubResponse struct {
	TopicId        string `json:"topicId"`
	SubscriptionId string `json:"subscriptionId"`
	ProjectId      string `json:"projectId"`
	ServiceAccount string `json:"serviceAccount"`
}

func (a *PubsubResponse) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
