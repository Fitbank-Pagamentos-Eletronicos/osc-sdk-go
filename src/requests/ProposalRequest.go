package requests

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osc-sdk-go/src/domains"
	"strings"
)

func ProposalRequest(token string, pipelineId string, proposalObject domains.ProposalReq) domains.Pipeline {
	url := "https://demo-api.easycredito.com.br/api/external/v2.1/process/proposal/" + pipelineId
	method := "POST"

	fmt.Println("URL:>", url)

	simpleProposalToJson, _ := json2.Marshal(proposalObject)
	payload := strings.NewReader(string(simpleProposalToJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return domains.Pipeline{}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

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
	json2.Unmarshal(body, &pipeline)
	return pipeline
}
