package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorField(t *testing.T) {
	errField := domains.ErrorField{
		Field:   "field",
		Message: "message",
	}

	assert.Equal(t, "field", errField.Field)
	assert.Equal(t, "message", errField.Message)

}
