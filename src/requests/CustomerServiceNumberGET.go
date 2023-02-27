package requests

import (
	json2 "encoding/json"
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"io/ioutil"
	"net/http"
)

func CustomerServiceNumberGET(token string) domains.GetContract {
	url := "https://demo-api.easycredito.com.br/api/external//v2.1/contract/20221109182327351003700"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return domains.GetContract{}
	}

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

	var contract domains.GetContract
	json2.Unmarshal(body, &contract)
	return contract

}
