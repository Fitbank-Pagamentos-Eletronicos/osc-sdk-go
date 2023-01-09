package main

import (
	"encoding/json"
	"fmt"
	"modulo/src/domains"
	"modulo/src/osc"
	"modulo/src/requests"
)

func main() {
	var response domains.PubsubResponse
	res := requests.PubSubRequest(&osc.OSC{})

	bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(bytes, &response)

	fmt.Println("==================Requisição PubSub==================")
	requests.PubSubSubscribe(
		response.Project_id,
		response.Topic_id,
		response.Subscription_id,
		response.Service_account,
		func(pipeline domains.Pipeline, success bool) {
			if success {
				fmt.Println("Pipeline: ", pipeline)
			} else {
				fmt.Println("Falha ao receber pipeline")
			}
		})

	fmt.Println("==================Fim da Requisição PubSub==================")

}
