package test

import (
	"osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHome(t *testing.T) {

	productHome := domains.ProductHome{
		Type:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		RealEstateType:     domains.HOUSE,
		RealEstateValue:    10000.00,
		OutstandingBalance: 1000.00,
	}

	assert.Equal(t, productHome.Type, domains.LOAN)
	assert.Equal(t, productHome.Value, 1000.00)
	assert.Equal(t, productHome.Installments, 12)
	assert.Equal(t, productHome.RealEstateType, domains.HOUSE)
	assert.Equal(t, productHome.RealEstateValue, 10000.00)
	assert.Equal(t, productHome.OutstandingBalance, 1000.00)

}
