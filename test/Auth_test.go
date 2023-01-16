package test

import (
	"github.com/stretchr/testify/assert"
	"osc-sdk-go/src/domains"
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
