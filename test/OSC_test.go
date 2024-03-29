package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/osc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOSC(t *testing.T) {
	var oscv osc.OSC
	getToken := domains.AuthSucess{
		Access_token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
		Expire_at:    "2022-11-26T10:56:24.038Z",
	}

	if getToken.Access_token == "" || getToken.Expire_at == "" {
		assert.Equal(t, getToken, oscv.GetToken())
		assert.Equal(t, getToken, oscv.GetToken())
	}

	normalized := osc.Normalize
	assert.Equal(t, normalized("@Pé de moleque"), "-pe-de-moleque")

}
