package requests

import (
	json2 "encoding/json"
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"io/ioutil"
	"net/http"
)

func PubSubRequest(token string) domains.PubsubResponse {
	url := "https://staging-api.easycredito.com.br/api/external/v2.1/pubsub"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return domains.PubsubResponse{}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

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
	json2.Unmarshal(body, &pubsubResponse)
	return pubsubResponse
}
