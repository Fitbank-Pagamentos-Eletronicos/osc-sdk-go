package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentResponse(t *testing.T) {
	documentResponse := domains.DocumentResponse{
		Type:     domains.IDENTITY_BACK,
		MimeType: domains.IMAGE_JPEG,
		Name:     "Teste",
		Url:      "Teste",
	}

	assert.Equal(t, domains.IDENTITY_BACK, documentResponse.Type)
	assert.Equal(t, domains.IMAGE_JPEG, documentResponse.MimeType)
	assert.Equal(t, "Teste", documentResponse.Name)
	assert.Equal(t, "Teste", documentResponse.Url)

}
