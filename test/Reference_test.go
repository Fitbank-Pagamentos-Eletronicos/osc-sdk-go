package test

import (
	"modulo/src/domains"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReference (t *testing.T) {
	reference := domains.Reference{
		Name:  "John Doe",
		Phone: "123456789",
	}

	assert.Equal(t, "John Doe", reference.Name)
	assert.Equal(t, "123456789", reference.Phone)
}