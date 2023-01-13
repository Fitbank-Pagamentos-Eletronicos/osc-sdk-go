package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthSucess(t *testing.T) {
	authSucess := domains.AuthSucess{
		Access_token: "teste",
		Expire_at:    "teste",
	}

	assert.Equal(t, "teste", authSucess.Access_token)
	assert.Equal(t, "teste", authSucess.Expire_at)
}
