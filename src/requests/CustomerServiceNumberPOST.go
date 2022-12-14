package requests


import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "modulo/src/domains"
    "encoding/json"

)
var baseContract = domains.Contract{
    AceptedCheckSum: [
        "97cc0c24610e982d38e2d28e80e7ff5af14bebd72491d548c1c5c1d2a4b7da06",
        "6cd99b452562c89d3cfccf2fd30c5e8633e59731795c89250be7d16cd3b034e1"
    ]

    LogData: domains.LogData{
        Latitude:  -16.6982283,
        Longitude: -49.2581201,
        OccurrenceDate: "2019-08-21T14:31:17.459Z",
        UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
        Ip: "1.0.0.0",
        Mac: "00.00.00.00.00.00"
    }
}

func CustomerServiceNumberPOST() string {
    url := "https://demo-api.easycredito.com.br/api/external//v2.1/contract/20221109182327351003700"
    method := "POST"

    simpleContractJson, _ := json.Marshal(baseContract)

    fmt.Println(string(simpleContractJson))

    payload := strings.NewReader(string(simpleContractJson))

    client := &http.Client{}

    req, err := http.NewRequest(method, url, payload)

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer " + GetToken())

    res, err := client.Do(req)
     if err != nil {
        fmt.Println(err)
        return ""
     }
     defer res.Body.Close()

     body, err := ioutil.ReadAll(res.Body)

        if err != nil {
           fmt.Println(err)
           return ""
        }
        fmt.Println(string(body))

        return string(body)

}