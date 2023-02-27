package requests

import (
	json2 "encoding/json"
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"io/ioutil"
	"net/http"
	"strings"
)

func SimpleProposalRequest(token string, ID string, dataSimpleProposal domains.ProposalReq) domains.ProposalReq {
	url := "https://demo-api.easycredito.com.br/api/external//v2.1/process/simple_proposal/" + ID
	method := "POST"

	fmt.Println("URL:>", url)

	jsonValue, _ := json2.Marshal(dataSimpleProposal)
	payload := strings.NewReader(string(jsonValue))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return domains.ProposalReq{}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domains.ProposalReq{}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return domains.ProposalReq{}
	}

	fmt.Println(string(body))
	var proposalBankAccount domains.ProposalReq
	json2.Unmarshal(body, &proposalBankAccount)
	return proposalBankAccount

}
