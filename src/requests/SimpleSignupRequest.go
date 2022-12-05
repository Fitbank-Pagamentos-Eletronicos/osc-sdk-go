package requests
import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "modulo/src/domains"
    "encoding/json"
)

var dataSimple = domains.SimpleSignup {

    Cpf: "05910188020",
    Name: "Renan Collienne Braga Alves",
    Birthday: "1992-05-07",
    Email: "renan@easyc.com.br",
    Phone: "62982103909",
    ZipCode: "74363820",

    LogData: domains.LogData{
         Latitude: -16.6982283,
         Longitude: -49.2581201,
         OccurrenceDate: "2019-08-21T14:31:17.459Z",
         UserAgent: "Mozilla/5.0 (Windows NT 10.0)",
         Ip: "0.0.0.0",
         Mac: "00:00:00:00:00:00",
     },
}

func SimpleSignupRequests() domains.SimpleSignup {

  url := "https://demo-api.easycredito.com.br/api/external//v2.1/process/simple_signup"
  method := "POST"
  signupToJson, _ := json.Marshal(dataSimple)
  payload := strings.NewReader(string(signupToJson))

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return domains.SimpleSignup{}
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Bearer " + GetToken())

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
  var simpleSignup domains.SimpleSignup
    json.Unmarshal(body, &simpleSignup)
    return simpleSignup

}