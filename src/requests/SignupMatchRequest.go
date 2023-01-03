package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modulo/src/domains"
	"modulo/src/osc"
	"net/http"
	"strings"
)

func SignupMatchRequest(osc *osc.OSC, signupObject domains.SignupMatch) domains.Pipeline {
	url := "https://demo-api.easycredito.com.br/api/external/v2.1/process/signup"
	method := "POST"

	simpleToJson, _ := json.Marshal(signupObject)

	payload := strings.NewReader(string(simpleToJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return domains.Pipeline{}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+osc.GetToken())

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
