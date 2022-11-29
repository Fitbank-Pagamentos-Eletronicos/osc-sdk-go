package requests

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "modulo/src/domains"
    "encoding/json"
)

func SimpleSignupRequest() domains.SimpleSignup {
    url:= "https://demo-api.easycredito.com.br/api/external//v2.1/process/simple_signup"
    method := "POST"

    payload := strings.NewReader(`{
        "cpf": "05910188020",
          "name": "Renan Collienne Braga Alves",
          "birthday": "1992-05-07",
          "email": "renan@easyc.com.br",
          "phone": "62982103909",
          "zipCode": "74363820",
          "logData": {
            "latitude": -16.6982283,
            "longitude": -49.2581201,
            "occurrenceDate": "2019-08-21T14:31:17.459Z",
            "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
            "ip": "0.0.0.0",
            "mac": "00:00:00:00:00:00"
          }

    }`)

    client := &http.Client {}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return domains.SimpleSignup{}
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEwNCwidGFnX25hbWUiOiI0ZGJlM2FhNy04Y2U5LTQzYTQtOTI5OC03M2I3MDBlNzEyYmIiLCJzY29wZSI6ImFwaS1leHRlcm5hbCIsImlhdCI6MTY2OTcyNTYzNiwiZXhwIjoxNjY5ODEyMDM2fQ.CwRoPncTlLis2fpZ1oz4ikKJ1nFwcFbfI9kCChYzUC1UNdAHLlbP_Ms_IwHInw2Nxpzuzy9WAoBp_1uAyzJYR4VfRSbHJdFxOPCUcLJMgn1-wyXZkeN64Ee-OyoUBWAwkRj1JhbSrqZXaXvyVcWA4oaa4HCn8KBeTp6CSdE2g7cs5mzDAdikCDksxoQ459Wwk7kJUxtQYYINWw5zIafXmBSMhHFO8cHhiW9Uiu2v-BbrBo6LP3dhRXBxXySuLENFNEu73I5PUIL1L8BJwQdcj22s_TLHUZrlbqeu6lnPcLob1AZJHarcaDHuW9w2X8HM3YV5VYLmHwamni7B7EMz_X2-u4XH4ovVSN1ajyuMthCxT5y8jaTQLculVyA8O53KCC4V75OCGTzcH3FhGyMvKFekzrHSEHXaXezzJHEgXS_T30xjLdkis4_V73wzF7z5Ya2Pt0STAkxu0Muz-BYg28dodRBEgaSwkhdGLUK3GVHB0pa-imnznlD6HXCGnw1pbBk3KvProcJv4puxjcjrfv9l3bT__txpD0pTUQl18q3gR5eGRfJ047xxToJTblM1yCFeGyviuFmMVwLB2WvoISnwOGjRZhmsFtTRFkaZ2OqyhAqwcWu-KnHA-qgW18LYgPxaV_FdTdhHX3ctdXu7k0baCzAec6WxRtIi3ZIlj6M")

    res, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        return domains.SimpleSignup{}
    }

    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Println(err)
        return domains.SimpleSignup{}
    }
    fmt.Println(string(body))

    var simple_signup domains.SimpleSignup
    json.Unmarshal(body, &simple_signup)

    return simple_signup

}