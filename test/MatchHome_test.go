package test



import (
	"testing"
	"modulo/src/domains"
	"github.com/stretchr/testify/assert"

)

func TestMatchHome(t *testing.T) {
	matchHome:= domains.MatchHome{
		ProductId:      1,
		Name:           "Casa",
		Logo:           "logo",
		MinValue:       1000,
		MaxValue:       10000,
		MinInstallment: 1,
		MaxInstallment: 12,
		MonthlyTax:     0.01,

	}
	assert.Equal(t, matchHome.ProductId, 1)
	assert.Equal(t, matchHome.Name, "Casa")
	assert.Equal(t, matchHome.Logo, "logo")
	assert.Equal(t, matchHome.MinValue, 1000)
	assert.Equal(t, matchHome.MaxValue, 10000)
	assert.Equal(t, matchHome.MinInstallment, 1)
	assert.Equal(t, matchHome.MaxInstallment, 12)
	assert.Equal(t, matchHome.MonthlyTax, 0.01)
}