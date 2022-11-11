package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchCard(t *testing.T) {

	matchCard := domains.MatchCard{
		ProductId: 2,
		Name:      "Cartão de Crédito",
		Logo:      "logo",
		Annuity:   1.00,
		Network:   domains.VISA,
	}

	fmt.Println(matchCard.ToJson())

	assert.Equal(t, 2, matchCard.ProductId)
	assert.Equal(t, "Cartão de Crédito", matchCard.Name)
	assert.Equal(t, "logo", matchCard.Logo)
	assert.Equal(t, 1.00, matchCard.Annuity)
	assert.Equal(t, domains.VISA, matchCard.Network)

}
