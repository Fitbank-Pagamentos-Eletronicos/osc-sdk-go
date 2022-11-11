package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchLoan(t *testing.T) {
	matchLoan := domains.MatchLoan{
		ProductId:      1,
		Name:           "Casa",
		Logo:           "logo",
		MinValue:       1000,
		MaxValue:       10000,
		MinInstallment: 1,
		MaxInstallment: 12,
		MonthlyTax:     0.01,
	}

	fmt.Println(matchLoan.ToJson())

	assert.Equal(t, matchLoan.ProductId, 1)
	assert.Equal(t, matchLoan.Name, "Casa")
	assert.Equal(t, matchLoan.Logo, "logo")
	assert.Equal(t, matchLoan.MinValue, 1000)
	assert.Equal(t, matchLoan.MaxValue, 10000)
	assert.Equal(t, matchLoan.MinInstallment, 1)
	assert.Equal(t, matchLoan.MaxInstallment, 12)
}
