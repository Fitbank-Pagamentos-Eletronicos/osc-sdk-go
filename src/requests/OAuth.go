package requests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modulo/src/domains"
	"net/http"
	"strings"
	"regexp"
)

var DataBase = domains.Auth{
	Client_id:     "4dbe3aa7-8ce9-43a4-9298-73b700e712bb",
	Client_secret: "1b364af124250aa09461f33161c3d96e551d822080fe1bd977aa66d7ec9378c8",
	Scopes:        []string{"api-external"},
}

func convertbase64(auth domains.Auth) string {
	return base64.StdEncoding.EncodeToString([]byte(auth.Client_id + ":" + auth.Client_secret))
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Expire_at   string `json:"expire_at"`
}

//Função responsavel por remover todos os caracteres especiais de uma string

func (a *AuthResponse) Normalize() string {
    var re = regexp.MustCompile("[^a-zA-Z0-9]+")
    return re.ReplaceAllString(a.AccessToken, "")
}

func (a *AuthResponse) OAuth() AuthResponse {
	url := "https://auth.easycredito.com.br/client/auth"
	method := "POST"
	payload := strings.NewReader(`{
      "scopes": ["api-external"]
  }`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return AuthResponse{}
	}
	req.Header.Add("Authorization", "Basic "+ convertbase64(DataBase))

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

	return authResponse
}
