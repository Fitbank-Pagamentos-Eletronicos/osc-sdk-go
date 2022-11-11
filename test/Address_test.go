package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress(t *testing.T) {
	address := domains.Address{
		ZipCode:    "00000000",
		Address:    "Rua Teste",
		Number:     "0",
		Complement: "Teste",
		District:   "Teste",
		State:      domains.ACRE,
		City:       "Teste",
		HomeType_:  domains.ALUGADA,
		HomeSince_: domains.MENOR_6_MESES,
	}

	fmt.Println(address.ToJson())

	assert.Equal(t, "00000000", address.ZipCode)
	assert.Equal(t, "Rua Teste", address.Address)
	assert.Equal(t, "0", address.Number)
	assert.Equal(t, "Teste", address.Complement)
	assert.Equal(t, "Teste", address.District)
	assert.Equal(t, domains.ACRE, address.State)
	assert.Equal(t, "Teste", address.City)
	assert.Equal(t, domains.ALUGADA, address.HomeType_)
	assert.Equal(t, domains.MENOR_6_MESES, address.HomeSince_)
}
