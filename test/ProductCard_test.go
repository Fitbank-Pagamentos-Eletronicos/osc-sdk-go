package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductCard(t *testing.T) {
	productCard := domains.ProductCard{
		Type:    domains.CARD,
		Network: domains.VISA,
		Payday:  "20",
	}
	fmt.Println(productCard.ToJson())

	assert.Equal(t, productCard.Type, domains.CARD)
	assert.Equal(t, productCard.Network, domains.VISA)
	assert.Equal(t, productCard.Payday, "20")
}
