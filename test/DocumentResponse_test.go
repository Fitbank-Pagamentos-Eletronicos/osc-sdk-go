package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentResponse(t *testing.T) {
	documentResponse := domains.DocumentResponse{
		Tipo:     domains.IDENTITY_BACK,
		MimeType: domains.IMAGE_JPEG,
		Name:     "Teste",
		Url:      "Teste",
	}

	fmt.Println(documentResponse.ToJson())

	assert.Equal(t, domains.IDENTITY_BACK, documentResponse.Tipo)
	assert.Equal(t, domains.IMAGE_JPEG, documentResponse.MimeType)
	assert.Equal(t, "Teste", documentResponse.Name)
	assert.Equal(t, "Teste", documentResponse.Url)

}
