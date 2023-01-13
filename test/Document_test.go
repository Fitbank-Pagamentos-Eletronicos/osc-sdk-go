package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	document := domains.Document{
		Type:     domains.IDENTITY_BACK,
		MimeType: domains.IMAGE_JPEG,
		Name:     "Teste",
		Base64:   "Teste",
	}

	fmt.Println(document.ToJson())

	assert.Equal(t, domains.IDENTITY_BACK, document.Type)
	assert.Equal(t, domains.IMAGE_JPEG, document.MimeType)
	assert.Equal(t, "Teste", document.Name)
	assert.Equal(t, "Teste", document.Base64)

}
