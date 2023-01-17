package requests

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

func CustomerServiceNumberPOST(token string, baseContract domains.Contract) domains.GetContract {
	url := "https://demo-api.easycredito.com.br/api/external//v2.1/contract/20221109182327351003700"
	method := "POST"

	simpleContractJson, _ := json2.Marshal(baseContract)

	payload := strings.NewReader(string(simpleContractJson))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domains.GetContract{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return domains.GetContract{}
	}
	fmt.Println(string(body))
	var getContract domains.GetContract
	json2.Unmarshal(body, &getContract)

	return getContract

}
