package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelineWithProposal(t *testing.T) {
	loan := domains.Loan{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.LOAN,
		Product:               "Empréstimo",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:	  domains.IDENTITY_BACK_,
		DateCreated:           "2020-01-01",
		LastUpdated:           "2020-01-01",
		Value:                 1000,
		Installments:          12,
		MonthlyTax:            10,
		InstallmentsValue:     100,
		IofValeu:              10,
		GrossValeu:            1000,
		FirstPaymentDate:      "2020-01-01",
		Cet:                   10,
		ReleaseDate:           "2020-01-01",
	}

	card := domains.Card{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.CARD,
		Product: 			 "Cartão de Crédito",
		ProductId:             2,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:	  domains.IDENTITY_BACK_,
		DateCreated:           "2020-01-01",
		LastUpdated:           "2020-01-01",
		International:         "Sim",
		Annuity:               "Sim",
		Network:               domains.VISA,
		Prepaid:               true,
		DigitalAccount:        true,
	}
	auto := domains.Auto{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.CARD,
		Product:               "Cartão de Crédito",
		ProductId:             2,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:	  domains.IDENTITY_BACK_,
		DateCreated:           "2020-01-01",
		LastUpdated:           "2020-01-01",
		Value:                 1000,
		Installments:          12,
		MonthlyTax:            10,
		InstallmentsValue:     100,
		IofValeu:              10,
		GrossValeu:            1000,
		FirstPaymentDate:      "2020-01-01",
		Cet:                   10,
		ReleaseDate:           "2020-01-01",
	}

	home := domains.Home{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.CARD,
		Product:               "Cartão de Crédito",
		ProductId:             2,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:	  domains.IDENTITY_BACK_,
		DateCreated:           "2020-01-01",
		LastUpdated:           "2020-01-01",
		Value:                 1000,
		Installments:          12,
		MonthlyTax:            10,
		InstallmentsValue:     100,
		IofValeu:              10,
		GrossValeu:            1000,
		FirstPaymentDate:      "2020-01-01",
		Cet:                   10,
		ReleaseDate:           "2020-01-01",
	}

	proposals := domains.Proposals{
		Loan: loan,
		Card: card,
		Auto: auto,
		Home: home,
	}

	pipelineWithProposal := domains.PipelineWithProposal{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "12345678901",
		Name:        "Felipe da Silva",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
		Proposals_:  proposals,
	}

	assert.Equal(t, "1", pipelineWithProposal.Id)
	assert.Equal(t, "Aprovado", pipelineWithProposal.Status)
	assert.Equal(t, "12345678901", pipelineWithProposal.Cpf)
	assert.Equal(t, "Felipe da Silva", pipelineWithProposal.Name)
	assert.Equal(t, "2020-01-01", pipelineWithProposal.DateCreated)
	assert.Equal(t, "2020-01-01", pipelineWithProposal.LastUpdated)
	assert.Equal(t, loan, pipelineWithProposal.Proposals_.Loan)
	assert.Equal(t, card, pipelineWithProposal.Proposals_.Card)
	assert.Equal(t, auto, pipelineWithProposal.Proposals_.Auto)
	assert.Equal(t, home, pipelineWithProposal.Proposals_.Home)
}
