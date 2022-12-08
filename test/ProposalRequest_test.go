package test

import (
    "modulo/src/requests"
    "modulo/src/domains"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestProposalRequest(t *testing.T)  {
    loan := domains.Loan{
    		CustomerServiceNumber: "123456789",
    		Tipo:                  domains.LOAN,
    		Product:               "Empréstimo",
    		ProductId:             1,
    		HasDocuments:          true,
    		HasContracts:          true,
    		Logo:                  "logo",
    		LastStatus:            domains.PRE_PROCESSAMENTO,
    		PendentDocuments:      domains.IDENTITY_BACK_,
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
    		Product:               "Cartão de Crédito",
    		ProductId:             2,
    		HasDocuments:          true,
    		HasContracts:          true,
    		Logo:                  "logo",
    		LastStatus:            domains.PRE_PROCESSAMENTO,
    		PendentDocuments:      domains.IDENTITY_BACK_,
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
    		PendentDocuments:      domains.IDENTITY_BACK_,
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
    		PendentDocuments:      domains.IDENTITY_BACK_,
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

    	assert.Equal(t, loan, proposals.Loan)
    	assert.Equal(t, card, proposals.Card)
    	assert.Equal(t, auto, proposals.Auto)
    	assert.Equal(t, home, proposals.Home)



    var id string = "1k8l2ugvfnpo16nff3i81d8d86e3b43443ba359b5c1c36bef23"

    _, status := requests.ProposalRequest(id)

        if status != 200 {
            assert.Equal(t, 200, status)
            t.Errorf("Erro ao fazer a requisição")
        } else {
            assert.Equal(t, 400, status)
            t.Log("Requisição feita com sucesso")
        }
}
