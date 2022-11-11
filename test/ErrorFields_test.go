package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFields(t *testing.T) {
	field := domains.ErrorField{
		Field:   "field",
		Message: "message",
	}

	fmt.Println(field.ToJson())

	errorFields := domains.ErrorFields{
		Code:    "code",
		Message: "message",
		Errors:  field,
	}
	fmt.Println(errorFields.ToJson())

	assert.Equal(t, "code", errorFields.Code)
	assert.Equal(t, "message", errorFields.Message)
	assert.Equal(t, field, errorFields.Errors)
}
