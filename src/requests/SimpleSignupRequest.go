package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

func SimpleSignupRequests(token string, signupObject domains.SimpleSignup) domains.Pipeline {

	url := "https://demo-api.easycredito.com.br/api/external/v2.1/process/simple_signup"
	method := "POST"
	signupToJson, _ := json.Marshal(signupObject)
	payload := strings.NewReader(string(signupToJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return domains.Pipeline{}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domains.Pipeline{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return domains.Pipeline{}
	}
	fmt.Println(string(body))
	var pipeline domains.Pipeline
	json.Unmarshal(body, &pipeline)
	return pipeline
}
