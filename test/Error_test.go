package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := domains.Error{
		Code:    "code",
		Message: "message",
	}

	fmt.Println(err.ToJson())

	assert.Equal(t, "code", err.Code)
	assert.Equal(t, "message", err.Message)

}
