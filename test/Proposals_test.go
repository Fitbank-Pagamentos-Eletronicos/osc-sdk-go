package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposals(t *testing.T) {
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

	fmt.Println(loan.ToJson())

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

	fmt.Println(card.ToJson())

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

	fmt.Println(auto.ToJson())

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

	fmt.Println(home.ToJson())

	proposals := domains.Proposals{
		Loan: loan,
		Card: card,
		Auto: auto,
		Home: home,
	}

	fmt.Println(proposals.ToJson())

	assert.Equal(t, loan, proposals.Loan)
	assert.Equal(t, card, proposals.Card)
	assert.Equal(t, auto, proposals.Auto)
	assert.Equal(t, home, proposals.Home)
}
