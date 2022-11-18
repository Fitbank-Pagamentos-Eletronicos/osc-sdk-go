package requests
import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "modulo/src/domains"
  "encoding/base64"
  "encoding/json"
  
)

var DataBase = domains.Auth{
	Client_id:     "4dbe3aa7-8ce9-43a4-9298-73b700e712bb",
	Client_secret: "1b364af124250aa09461f33161c3d96e551d822080fe1bd977aa66d7ec9378c8",
	Scopes:        []string{"api-external"},
}

func convertTobase64(auth domains.Auth) string {
  return base64.StdEncoding.EncodeToString([]byte(auth.Client_id + ":" + auth.Client_secret))
}

type AuthResponse struct {
  AccessToken string `json:"access_token"`
  Expire_at   string  `json:"expire_at"`
}


func OAuth() AuthResponse {
    url := "https://auth.easycredito.com.br/client/auth"
    method := "POST"
    payload := strings.NewReader(`{
      "scopes": ["api-external"]
  }`)
    client := &http.Client {
    }
    req, err := http.NewRequest(method, url, payload)
    if err != nil {
      fmt.Println(err)
      return AuthResponse{}
    }
    req.Header.Add("Authorization", "Basic " + convertTobase64(DataBase))
    
    req.Header.Add("Content-Type", "application/json")
    res, err := client.Do(req)
    if err != nil {
      fmt.Println(err)
      return AuthResponse{}
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
      fmt.Println(err)
      return AuthResponse{}
    }
    fmt.Println(string(body))
  
    var authResponse AuthResponse
  
    json.Unmarshal(body, &authResponse)
    
  
    fmt.Println(res.StatusCode)
    return authResponse
  
  }








