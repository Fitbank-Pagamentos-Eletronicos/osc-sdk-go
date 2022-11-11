package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankAccount(t *testing.T) {
	bankAccount := domains.BankAccount{
		CustomerServiceNumber: "123456789",
		Product:               "Conta Corrente",
		ProductId:             2,
		HasDocuments:          true,
		HasContracts:          true,
		LastStatus:            domains.PRE_PROCESSAMENTO,
		DateCreated:           "2020-01-01",
		LastUpdated:           "2020-01-01",
	}

	fmt.Println(bankAccount.ToJson())

	assert.Equal(t, "123456789", bankAccount.CustomerServiceNumber)
	assert.Equal(t, "Conta Corrente", bankAccount.Product)
	assert.Equal(t, 2, bankAccount.ProductId)
	assert.Equal(t, true, bankAccount.HasDocuments)
	assert.Equal(t, true, bankAccount.HasContracts)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, bankAccount.LastStatus)
	assert.Equal(t, "2020-01-01", bankAccount.DateCreated)
	assert.Equal(t, "2020-01-01", bankAccount.LastUpdated)

}
