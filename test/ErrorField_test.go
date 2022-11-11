package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorField(t *testing.T) {
	errField := domains.ErrorField{
		Field:   "field",
		Message: "message",
	}

	fmt.Println(errField.ToJson())

	assert.Equal(t, "field", errField.Field)
	assert.Equal(t, "message", errField.Message)

}
