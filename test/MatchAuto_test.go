package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchAuto(t *testing.T) {
	matchAuto := domains.MatchAuto{
		ProductId:      1,
		Name:           "Empréstimo",
		Logo:           "logo",
		MinValue:       1000,
		MaxValue:       10000,
		MinInstallment: 12,
		MaxInstallment: 24,
		MonthlyTax:     1.00,
	}

	fmt.Println(matchAuto.ToJson())

	assert.Equal(t, 1, matchAuto.ProductId)
	assert.Equal(t, "Empréstimo", matchAuto.Name)
	assert.Equal(t, "logo", matchAuto.Logo)
	assert.Equal(t, 1000, matchAuto.MinValue)
	assert.Equal(t, 10000, matchAuto.MaxValue)
	assert.Equal(t, 12, matchAuto.MinInstallment)
	assert.Equal(t, 24, matchAuto.MaxInstallment)
	assert.Equal(t, 1.00, matchAuto.MonthlyTax)
}
