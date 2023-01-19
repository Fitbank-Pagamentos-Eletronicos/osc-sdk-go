package requests

import (
	json2 "encoding/json"
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"io/ioutil"
	"net/http"
	"strings"
)

func CustomerServiceNumberPOST(token string, baseContract domains.Contract) domains.SignContract {
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
		return domains.SignContract{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return domains.SignContract{}
	}
	fmt.Println(string(body))
	var getContract domains.SignContract
	json2.Unmarshal(body, &getContract)

	return getContract

}
