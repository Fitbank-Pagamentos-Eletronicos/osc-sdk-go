package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

func OAuth(osc *OSC) domains.AuthSucess {
	url := "https://auth.easycredito.com.br/client/auth"

	method := "POST"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + osc.ClientId + "&client_secret=" + osc.ClientSecret + "&scope=api-external")

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

	return authSucess

}
