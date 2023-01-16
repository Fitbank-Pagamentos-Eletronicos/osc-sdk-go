package test

import (
	"osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductLoan(t *testing.T) {

	productLoan := domains.ProductLoan{
		Type:         domains.LOAN,
		Value:        1000,
		Installments: 12,
	}

	assert.Equal(t, productLoan.Type, domains.LOAN)
	assert.Equal(t, productLoan.Value, 1000)
	assert.Equal(t, productLoan.Installments, 12)

}
