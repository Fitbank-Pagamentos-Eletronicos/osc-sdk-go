package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelineWithBankAccountProposal(t *testing.T) {

	bankAccount := domains.BankAccount{
		CustomerServiceNumber: "123456789",
		Product:               "Conta Corrente",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		LastStatus:            domains.PRE_PROCESSAMENTO,
	}
	fmt.Println(bankAccount.ToJson())

	pipelineWithBankAccountProposal := domains.PipelineWithBankAccountProposal{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "12345678901",
		Name:        "Carla",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
		Proposals:   bankAccount,
	}

	assert.Equal(t, "1", pipelineWithBankAccountProposal.Id)
	assert.Equal(t, "Aprovado", pipelineWithBankAccountProposal.Status)
	assert.Equal(t, "12345678901", pipelineWithBankAccountProposal.Cpf)
	assert.Equal(t, "Carla", pipelineWithBankAccountProposal.Name)
	assert.Equal(t, "2020-01-01", pipelineWithBankAccountProposal.DateCreated)
	assert.Equal(t, "2020-01-01", pipelineWithBankAccountProposal.LastUpdated)
	assert.Equal(t, bankAccount, pipelineWithBankAccountProposal.Proposals)
}
