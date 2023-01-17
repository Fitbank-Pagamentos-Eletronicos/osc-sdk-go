package osc

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"osc-sdk-go/src/domains"
	"osc-sdk-go/src/requests"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type OSC struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
	AccessToken  string
	ExpireAt     time.Time
}

var instances map[string]OSC

func Normalize(name string) string {
	toConvert := regexp.MustCompile("[^A-z0-9]+")
	return toConvert.ReplaceAllString(strings.ToLower(strings.Trim(RemoveAccents(name), " ")), "-")
}

func RemoveAccents(str string) string {
	//goland:noinspection ALL
	format := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(format, str)
	return result
}

func isMn(runner rune) bool {
	return unicode.Is(unicode.Mn, runner)
}

func Auth(osc *OSC) domains.AuthSucess {
	return requests.OAuth(osc.ClientId, osc.ClientSecret)
}

func (osc *OSC) GetToken() string {
	if osc.AccessToken == "" || time.Time.IsZero(osc.ExpireAt) || time.Now().After(osc.ExpireAt) {
		var resp = Auth(osc)
		osc.AccessToken = resp.Access_token
		var ExpireAt, _ = time.Parse("2006-01-02T15:04:05.000Z", resp.Expire_at)
		osc.ExpireAt = ExpireAt
	}
	return osc.AccessToken
}

func CreateInstance(clientId string, clientSecret string, name string) (OSC, bool) {
	if instances == nil {
		instances = map[string]OSC{}
	}
	var normalized = Normalize(name)
	osc, ok := instances[normalized]
	if ok {
		if osc.ClientId == clientId && osc.ClientSecret == clientSecret {
			return osc, true
		}
		return OSC{}, false
	}

	osc = OSC{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"api-external"},
	}
	instances[normalized] = osc
	return osc, true
}

func GetInstance(name string) (OSC, bool) {
	if instances == nil {
		instances = map[string]OSC{}
	}
	var normalized = Normalize(name)
	osc, ok := instances[normalized]
	return osc, ok
}

func (osc *OSC) SetResponseListening(listeningFunction func(domains.Pipeline, bool)) bool {
	pubSubConfig := requests.PubSubRequest(osc.GetToken())
	return requests.PubSubSubscribe(pubSubConfig.Project_id, pubSubConfig.Topic_id, pubSubConfig.Subscription_id,
		pubSubConfig.Service_account, listeningFunction)
}

func (osc *OSC) SignupMatch(signupObject domains.SignupMatch) domains.Pipeline {
	pipeline := requests.SignupMatchRequest(osc.GetToken(), signupObject)
	return pipeline
}

func (osc *OSC) ContractGET() domains.GetContract {
	return requests.CustomerServiceNumberGET(osc.GetToken())
}

func (osc *OSC) ContractPOST(baseContract domains.Contract) domains.SignContract {
	return requests.CustomerServiceNumberPOST(osc.GetToken(), baseContract)
}

func (osc *OSC) SimpleSignup(signupObject domains.SimpleSignup) domains.Pipeline {
	pipeline := requests.SimpleSignupRequests(osc.GetToken(), signupObject)
	return pipeline
}

func (osc *OSC) SendDocument(id string, dataDocument domains.Document) domains.DocumentResponse {
	return requests.DocumentRequest(osc.GetToken(), id, dataDocument)
}

func (osc *OSC) Proposal(pipelineId string, proposalObject domains.ProposalReq) domains.Pipeline {
	pipeline := requests.ProposalRequest(osc.GetToken(), pipelineId, proposalObject)
	return pipeline
}

func (osc *OSC) SimpleProposal(id string, dataSimpleProposal domains.ProposalReq) domains.ProposalReq {
	return requests.SimpleProposalRequest(osc.GetToken(), id, dataSimpleProposal)
}
