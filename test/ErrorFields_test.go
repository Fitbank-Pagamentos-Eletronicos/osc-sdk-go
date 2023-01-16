package test

import (
	"osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFields(t *testing.T) {
	field := domains.ErrorField{
		Field:   "field",
		Message: "message",
	}

	errorFields := domains.ErrorFields{
		Code:    "code",
		Message: "message",
		Errors:  field,
	}

	assert.Equal(t, "code", errorFields.Code)
	assert.Equal(t, "message", errorFields.Message)
	assert.Equal(t, field, errorFields.Errors)
}
