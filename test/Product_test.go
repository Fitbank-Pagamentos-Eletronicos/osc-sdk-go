package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	product := domains.Product{
		ProductId: 1,
		Name:      "Empréstimo",
		Logo:      "logo",
	}

	assert.Equal(t, product.ProductId, 1)
	assert.Equal(t, product.Name, "Empréstimo")
	assert.Equal(t, product.Logo, "logo")
}
