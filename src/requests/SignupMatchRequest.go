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
        Cpf: "504.768.940-93",
        Name: "Fulano d'tal",
        Birthday: "1992-07-25",
        Email: "fulano6243242@email.com",
        Phone: "62987832321",
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
            Latitude: -16.6982283,
            Longitude: -49.2581201,
            OccurrenceDate: "2019-08-21T14:31:17.459Z",
            UserAgent: "Mozilla/5.0 (Windows NT 10.0)",
            Ip: "0.0.0.0",
            Mac: "00:00:00:00:00:00",
        },

}


func SignupMatchRequest() domains.SignupMatch {
    url:= "https://demo-api.easycredito.com.br/api/external//v2.1/process/signup"
    method := "POST"

    simpleToJson, _ := json.Marshal(data)
    //fmt.Println("Função Marshal", string(simpleToJson))
    payload := strings.NewReader(string(simpleToJson))

    client := &http.Client{}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return domains.SignupMatch{}
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEwNCwidGFnX25hbWUiOiI0ZGJlM2FhNy04Y2U5LTQzYTQtOTI5OC03M2I3MDBlNzEyYmIiLCJzY29wZSI6ImFwaS1leHRlcm5hbCIsImlhdCI6MTY3MDI0MjAwMSwiZXhwIjoxNjcwMzI4NDAxfQ.qOor64qT_JrA9hiFpk8HC4PcFl_CXbNDF_bqSLaEnArzKFOvUv8jNG-AMacFGBKj9WqZXhTB5PeG7RHD-_sMevGW7O-_XZdxrNyaZ5Qtibx-oCm7y5ztceH7k8Ll74pY3PnmQ4LhygiIsiagOWqD7MNJSgeLi10fP0N8sP-bY9RSl2QachweD_SCcvaI60E5tTM-SpB4xDDXyVOZKyk_e5EvRorJgGxZaPsrSIegevbnp6SpZRhyc9_WLbw165XsKRXlKmC6L6wlmn2MzJVeNBdN-_BCNIC4Ad5PeO8YlwZFTxJ0Jqs1sUmRs0pZwL-FxBYiEDJ1Zbi5DPmesO2Zh4eys5AzBPZlQiiBHffvXoKnjfozkF0MFYP-2U4jbHoTFgGox4wGBhNCvFCMQ5MqVJxR5zJv1c2LpVEr8T3SFqiH2pU7qMpQ-BUuGOOVshAjR0byHkuBHaCCZtAeTbFeH1bhyk-PKZUX5lYtyT2vPlbnE40-cFgsDWyXapHROu2yeEmlBnQFGZN-OstcRGWh3g2tPARgtyRMomD-Syx9YpYsDlO6CqNMf3icDh94sBF3DxSqYNp-IIA2VtNoDqASBRrlEpcH6jHQEan6NbtYV84LCop_-WaBmnmoYpTslbOilUEFenF0bGIPi-np2vqAtu6aE95OnT3USFiMxP_638o" )

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return domains.SignupMatch{}
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Println(err)
        return domains.SignupMatch{}
    }

    fmt.Println(string(body))

    var signupMatch domains.SignupMatch
    json.Unmarshal(body, &signupMatch)
    return signupMatch
}