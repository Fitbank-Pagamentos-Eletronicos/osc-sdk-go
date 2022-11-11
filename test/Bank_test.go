package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank(t *testing.T) {
	bank := domains.Bank{
		Bank_:   domains.BANCO_DO_BRASIL,
		Tipo:    domains.CONTA_CORRENTE_INDIVIDUAL,
		Agency:  "0001",
		Account: "123456",
	}

	fmt.Println(bank.ToJson())

	assert.Equal(t, domains.BANCO_DO_BRASIL, bank.Bank_)
	assert.Equal(t, domains.CONTA_CORRENTE_INDIVIDUAL, bank.Tipo)
	assert.Equal(t, "0001", bank.Agency)
	assert.Equal(t, "123456", bank.Account)
}
