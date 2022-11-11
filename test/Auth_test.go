package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	auth := domains.Auth{
		Client_id:     "teste",
		Client_secret: "teste",
		Scopes:        []string{"teste"},
	}

	fmt.Println(auth.ToJson())

	assert.Equal(t, "teste", auth.Client_id)
	assert.Equal(t, "teste", auth.Client_secret)
	assert.Equal(t, []string{"teste"}, auth.Scopes)
}
