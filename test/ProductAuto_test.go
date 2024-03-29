package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductAuto(t *testing.T) {
	productAuto := domains.ProductAuto{
		Type:             domains.LOAN,
		Value:            1000.00,
		VehicleBrand:     "Fiat",
		VehicleModel:     "Uno",
		Installments:     12,
		CodeFipe:         "123",
		VehicleFipeValue: 10000.00,
	}

	assert.Equal(t, productAuto.Type, domains.LOAN)
	assert.Equal(t, productAuto.Value, 1000.00)
	assert.Equal(t, productAuto.VehicleBrand, "Fiat")
	assert.Equal(t, productAuto.VehicleModel, "Uno")
	assert.Equal(t, productAuto.Installments, 12)
	assert.Equal(t, productAuto.CodeFipe, "123")
	assert.Equal(t, productAuto.VehicleFipeValue, 10000.00)

}
