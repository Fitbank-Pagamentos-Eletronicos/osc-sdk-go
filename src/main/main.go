package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/requests"
)

func main() {
	var response domains.PubsubResponse
	res := requests.PubSubRequest(&requests.OSC{})

	bytes, err := json2.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	json2.Unmarshal(bytes, &response)

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
