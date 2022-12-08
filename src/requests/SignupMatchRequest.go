package requests

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
    "modulo/src/domains"
)

var data = domains.SignupMatch {
        Cpf: "720.825.560-18",
        Name: "Carlos Henrique",
        Birthday: "2002-01-19",
        Email: "CH624323242@email.com",
        Phone: "63980030021",
        ZipCode : "74180040",
        Education: domains.POS_GRADUACAO,
        Banks: domains.BANCO_DO_BRASIL,
        Occupation: domains.ASSALARIADO,
        Income: 1045,
        HasCreditCart: true,
        HasRestriction: true,
        HasOwnHouse: true,
        HasVehicle: true,
        HasAndroid: true,

        Products: domains.Products{

            ProductLoan: domains.ProductLoan{
               Tipo: domains.LOAN,
               Value: 7000,
               Installments: 12,
            },

            ProductCard: domains.ProductCard {
                Tipo: domains.CARD,
                Network: domains.MASTERCARD,
                Payday: "15",
            },

            ProductAuto: domains.ProductAuto{
                Tipo: domains.REFINANCING_AUTO,
                Value: 3000.00,
                VehicleBrand: "Fiat",
                VehicleModel: "Mobi",
                Installments: 12,
                VehicleModelYear: "2010",
                CodeFipe : "038003-2" ,
                VehicleFipeValue: 28000.00,
            },

            ProductHome: domains.ProductHome{
                Tipo: domains.REFINANCING_HOME,
                Value: 15000.00,
                Installments: 12,
                RealEstateType_: domains.HOUSE,
                RealEstateValue: 148000.00,
                OutstandingBalance: 50000.00,
             },

        },

        LogData: domains.LogData{
            Latitude: -16.678,
            Longitude: -49.256,
            OccurrenceDate: "2019-08-21T14:31:17.459Z",
            UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36",
            Ip: "0.0.0.0",
            Mac: "00:00:00:00:00:00",

        },

}

func SignupMatchRequest() string {
    url:= "https://demo-api.easycredito.com.br/api/external//v2.1/process/signup"
    method := "POST"

    simpleToJson, _ := json.Marshal(data)

    payload := strings.NewReader(string(simpleToJson))

    client := &http.Client{}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return  ""
    }

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

    return string(body)
}
