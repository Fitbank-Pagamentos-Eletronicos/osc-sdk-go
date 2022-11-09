package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthSucess(t *testing.T) {
	authSucess := domains.AuthSucess{
		Acess_token: "teste",
		Expire_at:   "teste",
	}

	assert.Equal(t, "teste", authSucess.Acess_token)
	assert.Equal(t, "teste", authSucess.Expire_at)
}
