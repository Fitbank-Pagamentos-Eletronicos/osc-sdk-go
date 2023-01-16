package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
)

func PubSubRequest(osc *OSC) domains.PubsubResponse {
	url := "https://staging-api.easycredito.com.br/api/external/v2.1/pubsub"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return domains.PubsubResponse{}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+osc.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domains.PubsubResponse{}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return domains.PubsubResponse{}
	}
	fmt.Println(string(body))
	var pubsubResponse domains.PubsubResponse
	json.Unmarshal(body, &pubsubResponse)
	return pubsubResponse
}
