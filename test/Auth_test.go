package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"modulo/src/domains"
	"testing"
)

func TestAuth(t *testing.T) {
	auth := domains.Auth{
		ClientId:     "teste",
		ClientSecret: "teste",
		Scopes:       []string{"teste"},
	}

	fmt.Println(auth.ToJson())

	assert.Equal(t, "teste", auth.ClientId)
	assert.Equal(t, "teste", auth.ClientSecret)
	assert.Equal(t, []string{"teste"}, auth.Scopes)
}
