package main

import (
    "modulo/src/requests"
    "modulo/src/domains"
    "encoding/json"
    "fmt"
    "time"
)

func main() {

    var SignupResponse domains.SignupMatchResponse
    response := requests.SignupMatchRequest()

    json.Unmarshal([]byte(response), &SignupResponse)

    time.Sleep(10 * time.Second)


    fmt.Println("==================Requisição de DocumentRequest==================")
    requests.DocumentRequest(SignupResponse.ID)



}
