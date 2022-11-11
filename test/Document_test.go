package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	document := domains.Document{
		Tipo:      domains.IDENTITY_BACK,
		MimeType_: domains.IMAGE_JPEG,
		Name:      "Teste",
		Base64:    "Teste",
	}

	fmt.Println(document.ToJson())

	assert.Equal(t, domains.IDENTITY_BACK, document.Tipo)
	assert.Equal(t, domains.IMAGE_JPEG, document.MimeType_)
	assert.Equal(t, "Teste", document.Name)
	assert.Equal(t, "Teste", document.Base64)

}
