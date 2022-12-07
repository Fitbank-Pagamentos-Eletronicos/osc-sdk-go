package requests

import (
	"encoding/base64"
	"encoding/json"
    "io/ioutil"
    "fmt"
	"modulo/src/domains"
	"net/http"
	"strings"
)

var DataBase = domains.Auth{
	Client_id:     "4dbe3aa7-8ce9-43a4-9298-73b700e712bb",
	Client_secret: "1b364af124250aa09461f33161c3d96e551d822080fe1bd977aa66d7ec9378c8",
	Scopes:        []string{"api-external"},
}

func convertbase64(auth domains.Auth) string {
	return base64.StdEncoding.EncodeToString([]byte(auth.Client_id + ":" + auth.Client_secret))
}

func OAuth() domains.AuthSucess {
	url := "https://auth.easycredito.com.br/client/auth"
       method := "POST"
       payload := strings.NewReader("grant_type=client_credentials&client_id=" + DataBase.Client_id + "&client_secret=" + DataBase.Client_secret + "&scope=api-external")

       client := &http.Client{}
       req, err := http.NewRequest(method, url, payload)

       if err != nil {
          fmt.Println(err)
          return domains.AuthSucess{}
       }
       req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

       res, err := client.Do(req)
       if err != nil {
          fmt.Println(err)
          return domains.AuthSucess{}
       }
       defer res.Body.Close()
       body, err := ioutil.ReadAll(res.Body)
       if err != nil {
          fmt.Println(err)
          return domains.AuthSucess{}
       }

       fmt.Println(string(body))
       var authSucess domains.AuthSucess
       json.Unmarshal(body, &authSucess)

       var response map[string]interface{}
       json.Unmarshal([]byte(body), &response)
       fmt.Println("===================================")
       fmt.Println(response["access_token"])


       return authSucess
}
