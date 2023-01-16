package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipelineWithProposal(t *testing.T) {
	proposal := domains.Proposal{
		CustomerServiceNumber: "123456789",
		Type:                  domains.LOAN,
		Product:               "2",
		ProductId:             12,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		DateCreated:           time.Now(),
		LastUpdated:           time.Now(),
		PendentDocuments:      domains.SELF_,
	}
	pipelineWithProposal := domains.PipelineWithProposal{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "12345678901",
		Name:        "Felipe da Silva",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
		Proposals:   proposal,
	}

	assert.Equal(t, "1", pipelineWithProposal.Id)
	assert.Equal(t, "Aprovado", pipelineWithProposal.Status)
	assert.Equal(t, "12345678901", pipelineWithProposal.Cpf)
	assert.Equal(t, "Felipe da Silva", pipelineWithProposal.Name)
	assert.Equal(t, "2020-01-01", pipelineWithProposal.DateCreated)
	assert.Equal(t, "2020-01-01", pipelineWithProposal.LastUpdated)
	assert.Equal(t, proposal, pipelineWithProposal.Proposals)
}
