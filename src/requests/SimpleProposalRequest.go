package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

var dataSimpleProposal = domains.ProposalBankAccount{
	Mother:             "Fulana Mãe",
	Gender:             domains.FEMININO,
	Nationality:        domains.BRASILEIRO,
	HomeTownState:      domains.GOIAS,
	RelationshipStatus: domains.CASADO,

	Address: domains.Address{
		ZipCode:    "74000000",
		Address:    "Cidade de Goiânia",
		Number:     "0",
		Complement: "Casa 1",
		District:   "Geral",
		State:      domains.GOIAS,
		City:       "Goiânia",
	},
	Business: domains.Business{
		Income: "1000.00",
	},
	Products: domains.ProductBankAccount{
		Type: "BANK_ACCOUNT",
	},
}

func SimpleProposalRequest(osc *OSC, ID string) (string, int) {
	url := "https://demo-api.easycredito.com.br/api/external//v2.1/process/simple_proposal/" + ID
	method := "POST"

	fmt.Println("URL:>", url)

	jsonValue, _ := json.Marshal(dataSimpleProposal)
	payload := strings.NewReader(string(jsonValue))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", 0
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+osc.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", 0
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", 0
	}

	fmt.Println(string(body))
	return string(body), res.StatusCode

}
