package main

import (
	json2 "encoding/json"
	"fmt"
	"osc-sdk-go/src/domains"
	"osc-sdk-go/src/requests"
)

func main() {
	var response domains.PubsubResponse
	res := requests.PubSubRequest("1346579")

	bytes, err := json2.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	json2.Unmarshal(bytes, &response)

	requests.PubSubSubscribe(response.ProjectId, response.TopicId, response.SubscriptionId, response.ServiceAccount,
		func(pipeline domains.Pipeline, success bool) {
			if success {
				fmt.Println("Pipeline: ", pipeline)
			} else {
				fmt.Println("Falha ao receber pipeline")
			}
		})

	fmt.Println("==================Fim da Requisição PubSub==================")

}
