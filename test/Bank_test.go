package test

import (
	"osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank(t *testing.T) {
	bank := domains.Bank{
		Bank:    domains.BANCO_DO_BRASIL,
		Type:    domains.CONTA_CORRENTE_INDIVIDUAL,
		Agency:  "0001",
		Account: "123456",
	}

	assert.Equal(t, domains.BANCO_DO_BRASIL, bank.Bank)
	assert.Equal(t, domains.CONTA_CORRENTE_INDIVIDUAL, bank.Type)
	assert.Equal(t, "0001", bank.Agency)
	assert.Equal(t, "123456", bank.Account)
}
