package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	loan := domains.ProductLoan{
		Tipo:         domains.LOAN,
		Value:        1000,
		Installments: 12,
	}
	card := domains.ProductCard{
		Tipo:    domains.CARD,
		Network: domains.VISA,
		Payday:  "20",
	}
	auto := domains.ProductAuto{
		Tipo:             domains.LOAN,
		Value:            1000.00,
		VehicleBrand:     "Fiat",
		VehicleModel:     "Uno",
		Installments:     12,
		CodeFipe:         "123",
		VehicleFipeValue: 10000.00,
	}

	home := domains.ProductHome{
		Tipo:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		RealEstateType_:    domains.HOUSE,
		RealEstateValue:    10000.00,
		OutstandingBalance: 1000.00,
	}
	products := domains.Products{
		ProductLoan: loan,
		ProductCard: card,
		ProductAuto: auto,
		ProductHome: home,
	}

	assert.Equal(t, loan, products.ProductLoan)
	assert.Equal(t, card, products.ProductCard)
	assert.Equal(t, auto, products.ProductAuto)
	assert.Equal(t, home, products.ProductHome)

}
