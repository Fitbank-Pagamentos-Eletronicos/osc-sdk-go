package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := domains.Error{
		Code:    "code",
		Message: "message",
	}

	assert.Equal(t, "code", err.Code)
	assert.Equal(t, "message", err.Message)

}
