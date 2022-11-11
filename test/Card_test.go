package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	card := domains.Card{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.CARD,
		Product:               "Cartão de Crédito",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:      domains.SELF_,
		DateCreated:           "01/01/2020",
		LastUpdated:           "01/01/2020",
		International:         "Sim",
		Annuity:               "Sim",
		Network:               domains.VISA,
		Prepaid:               true,
		DigitalAccount:        true,
	}

	fmt.Println(card.ToJson())

	assert.Equal(t, "123456789", card.CustomerServiceNumber)
	assert.Equal(t, domains.CARD, card.Tipo)
	assert.Equal(t, "Cartão de Crédito", card.Product)
	assert.Equal(t, 1, card.ProductId)
	assert.Equal(t, true, card.HasDocuments)
	assert.Equal(t, true, card.HasContracts)
	assert.Equal(t, "logo", card.Logo)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, card.LastStatus)
	assert.Equal(t, domains.SELF_, card.PendentDocuments)
	assert.Equal(t, "01/01/2020", card.DateCreated)
	assert.Equal(t, "01/01/2020", card.LastUpdated)
	assert.Equal(t, "Sim", card.International)
	assert.Equal(t, "Sim", card.Annuity)
	assert.Equal(t, domains.VISA, card.Network)
	assert.Equal(t, true, card.Prepaid)
	assert.Equal(t, true, card.DigitalAccount)
}
