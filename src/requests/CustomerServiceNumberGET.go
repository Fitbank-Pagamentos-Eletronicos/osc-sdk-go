package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CustomerServiceNumberGET(osc *OSC) string {
	url := "https://demo-api.easycredito.com.br/api/external//v2.1/contract/20221109182327351003700"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+osc.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(body))

	return string(body)

}
