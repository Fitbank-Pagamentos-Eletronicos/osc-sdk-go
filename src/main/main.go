package main

import (
    "modulo/src/requests"
    "encoding/json"
    "fmt"

)
func main() {
     var response requests.PubsubResponse
     res := requests.PubSubRequest()

     json.Unmarshal([]byte(res), &response)

    fmt.Println("==================Requisição PubSub==================")
    requests.PubSubSubscribe(response.Project_id, response.Topic_id, response.Subscription_id, response.Service_account)
    fmt.Println("==================Fim da Requisição PubSub==================")

}
