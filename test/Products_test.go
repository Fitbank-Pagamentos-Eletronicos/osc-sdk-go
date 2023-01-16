package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	products := domains.Products{
		Type:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		Network:            domains.VISA,
		Payday:             "20",
		VehicleBrand:       "Fiat",
		VehicleModel:       "Uno",
		CodeFipe:           "123",
		VehicleFipeValue:   10000.00,
		RealEstateType:     domains.HOUSE,
		RealEstateValue:    10000.00,
		OutstandingBalance: 1000.00,
	}

	assert.Equal(t, products.Type, domains.LOAN)
	assert.Equal(t, products.Value, 1000.00)
	assert.Equal(t, products.Installments, 12)
	assert.Equal(t, products.Network, domains.VISA)
	assert.Equal(t, products.Payday, "20")
	assert.Equal(t, products.VehicleBrand, "Fiat")
	assert.Equal(t, products.VehicleModel, "Uno")
	assert.Equal(t, products.CodeFipe, "123")
	assert.Equal(t, products.VehicleFipeValue, 10000.00)
	assert.Equal(t, products.RealEstateType, domains.HOUSE)
	assert.Equal(t, products.RealEstateValue, 10000.00)
	assert.Equal(t, products.OutstandingBalance, 1000.00)

}
