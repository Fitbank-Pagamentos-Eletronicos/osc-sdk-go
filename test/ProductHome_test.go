package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHome(t *testing.T) {

	productHome := domains.ProductHome{
		Tipo:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		RealEstateType_:    domains.HOUSE,
		RealEstateValue:    10000.00,
		OutstandingBalance: 1000.00,
	}

	fmt.Println(productHome.ToJson())

	assert.Equal(t, productHome.Tipo, domains.LOAN)
	assert.Equal(t, productHome.Value, 1000.00)
	assert.Equal(t, productHome.Installments, 12)
	assert.Equal(t, productHome.RealEstateType_, domains.HOUSE)
	assert.Equal(t, productHome.RealEstateValue, 10000.00)
	assert.Equal(t, productHome.OutstandingBalance, 1000.00)

}
