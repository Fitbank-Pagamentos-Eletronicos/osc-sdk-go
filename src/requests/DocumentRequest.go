package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

func DocumentRequest(token string, id string, dataDocument domains.Document) domains.DocumentResponse {
	url := "https://demo-api.easycredito.com.br/api/external/v2/process/document/" + id
	method := "PUT"

	fmt.Println("URL: ", url)
	simpleDocumentJson, _ := json.Marshal(dataDocument)

	payload := strings.NewReader(string(simpleDocumentJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return domains.DocumentResponse{}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return domains.DocumentResponse{}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return domains.DocumentResponse{}
	}

	fmt.Println(string(body))

	return string(body)

}
