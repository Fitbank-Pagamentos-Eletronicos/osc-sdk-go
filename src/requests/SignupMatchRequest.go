package requests

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "modulo/src/domains"
    "encoding/json"
)

func SignupMatchRequest() domains.SignupMatch {
    url:= "https://demo-api.easycredito.com.br/api/external//v2.1/process/signup"
    method := "POST"

    payload := strings.NewReader(`{
      "cpf": "504.768.940-93",
      "name": "Carlos Alexandre",
      "birthday": "1992-07-25",
      "email": "fulano6243242@email.com",
      "phone": "62987832321",
      "zipCode": "74180040",
      "education": "POS_GRADUACAO",
      "banks": "001,033",
      "occupation": "ASSALARIADO",
      "income": 1045,
      "hasCreditCard": true,
      "hasRestriction": true,
      "hasOwnHouse": true,
      "hasVehicle": true,
      "hasAndroid": true,
      "products": [
        {
          "type": "LOAN",
          "value": "7000",
          "installments": 12
        },
        {
          "type": "CARD",
          "network": "MASTERCARD",
          "payDay": 15
        },
        {
          "type": "REFINANCING_AUTO",
          "value": "30000",
          "vehicleBrand": "Fiat",
          "vehicleModel": "Mobi",
          "installments": 12,
          "vehicleModelYear": "2010",
          "codeFipe": "038003-2",
          "vehicleFipeValue": "28000"
        },
        {
          "type": "REFINANCING_HOME",
          "value": "150000",
          "installments": 12,
          "realEstateType": "house",
          "realEstateValue": "148000",
          "outstandingBalance": "50000"
        }
      ],
      "logData": {
        "latitude": -16.6982283,
        "longitude": -49.2581201,
        "occurrenceDate": "2019-08-21T14:31:17.459Z",
        "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
        "ip": "0.0.0.0",
        "mac": "00:00:00:00:00:00"
      }
    }`)

    client := &http.Client{}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return domains.SignupMatch{}
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEwNCwidGFnX25hbWUiOiI0ZGJlM2FhNy04Y2U5LTQzYTQtOTI5OC03M2I3MDBlNzEyYmIiLCJzY29wZSI6ImFwaS1leHRlcm5hbCIsImlhdCI6MTY2OTcyNTYzNiwiZXhwIjoxNjY5ODEyMDM2fQ.CwRoPncTlLis2fpZ1oz4ikKJ1nFwcFbfI9kCChYzUC1UNdAHLlbP_Ms_IwHInw2Nxpzuzy9WAoBp_1uAyzJYR4VfRSbHJdFxOPCUcLJMgn1-wyXZkeN64Ee-OyoUBWAwkRj1JhbSrqZXaXvyVcWA4oaa4HCn8KBeTp6CSdE2g7cs5mzDAdikCDksxoQ459Wwk7kJUxtQYYINWw5zIafXmBSMhHFO8cHhiW9Uiu2v-BbrBo6LP3dhRXBxXySuLENFNEu73I5PUIL1L8BJwQdcj22s_TLHUZrlbqeu6lnPcLob1AZJHarcaDHuW9w2X8HM3YV5VYLmHwamni7B7EMz_X2-u4XH4ovVSN1ajyuMthCxT5y8jaTQLculVyA8O53KCC4V75OCGTzcH3FhGyMvKFekzrHSEHXaXezzJHEgXS_T30xjLdkis4_V73wzF7z5Ya2Pt0STAkxu0Muz-BYg28dodRBEgaSwkhdGLUK3GVHB0pa-imnznlD6HXCGnw1pbBk3KvProcJv4puxjcjrfv9l3bT__txpD0pTUQl18q3gR5eGRfJ047xxToJTblM1yCFeGyviuFmMVwLB2WvoISnwOGjRZhmsFtTRFkaZ2OqyhAqwcWu-KnHA-qgW18LYgPxaV_FdTdhHX3ctdXu7k0baCzAec6WxRtIi3ZIlj6M")

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