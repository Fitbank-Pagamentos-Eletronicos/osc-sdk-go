package main

import (
    "modulo/src/requests"
    "modulo/src/domains"
    "encoding/json"
    "fmt"
    "time"
)

func main() {
    requests.SignupMatchRequest()

    var SignupResponse domains.SignupMatchResponse
    response := requests.SignupMatchRequest()

    json.Unmarshal([]byte(response), &SignupResponse)

    time.Sleep(10 * time.Second)

   // fmt.Println("==================Requisição de ProposalRequest==================")
   // requests.ProposalRequest(SignupResponse.ID)


    fmt.Println("==================Requisição de SimpleProposalRequest==================")
    requests.SimpleProposalRequest(SignupResponse.ID)

}
