package requests

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
)

func ProposalRequest() domains.Proposal{
   url := "https://demo-api.easycredito.com.br/api/external//v2.1/process/proposal/1k8l2p52h7donhqa9sdae55657ffe364e6d8b56747bc434035e"
   method := "POST"

   payload := strings.NewReader()


     client := &http.Client {}
     req, err := http.NewRequest(method, url, payload)

     if err != nil {
       fmt.Println(err)
       return
     }
     req.Header.Add("Content-Type", "application/json")
     req.Header.Add("Accept", "application/json")
     req.Header.Add("Authorization", "Bearer " + GetToken())

     res, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
            return domains.Proposal{}
        }

        defer res.Body.Close()

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println(err)
            return domains.Proposal{}
        }
        fmt.Println(string(body))

        var proposal domains.Proposal
        json.Unmarshal(body, &proposal)
        return proposal

}