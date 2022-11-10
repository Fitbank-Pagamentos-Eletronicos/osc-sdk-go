package test


import (
	"testing"
	"modulo/src/domains"
	"github.com/stretchr/testify/assert"
)


func TestProductCard(t *testing.T) {
	productCard := domains.ProductCard{
		Tipo:    domains.CARD,
		Network: domains.VISA,
		Payday:  "20",
	}

	assert.Equal(t, productCard.Tipo, domains.CARD)
	assert.Equal(t, productCard.Network, domains.VISA)
	assert.Equal(t, productCard.Payday, "20")
}