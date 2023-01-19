package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	auth := domains.Auth{
		ClientId:     "teste",
		ClientSecret: "teste",
		Scopes:       []string{"teste"},
	}

	assert.Equal(t, "teste", auth.ClientId)
	assert.Equal(t, "teste", auth.ClientSecret)
	assert.Equal(t, []string{"teste"}, auth.Scopes)
}
