package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductLoan(t *testing.T) {

	productLoan := domains.ProductLoan{
		Tipo:         domains.LOAN,
		Value:        1000,
		Installments: 12,
	}

	fmt.Println(productLoan.ToJson())

	assert.Equal(t, productLoan.Tipo, domains.LOAN)
	assert.Equal(t, productLoan.Value, 1000)
	assert.Equal(t, productLoan.Installments, 12)

}
