package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposalBankAccount(t *testing.T) {
	address := domains.Address{
		ZipCode:    "12345678",
		Address:    "Rua dos Bobos",
		Number:     "0",
		Complement: "Casa",
		District:   "Bairro",
		State:      domains.SANTA_CANTARINA,
		City:       "Florianópolis",
	}
	fmt.Println(address.ToJson())

	business := domains.Business{
		Occupation_:      domains.ASSALARIOADO,
		Profession_:      domains.ACUMPURISTA,
		CompanyName:      "Empresa",
		Income:           "1000",
		EmploymentSince_: domains.LESS_THAN_SIX_MONTHS,
		Payday:           "10",
		BenefitNumber:    "123456789",
		ZipCode:          "12345678",
		Address:          "Rua dos Bobos",
		Number:           "0",
		Complement:       "Casa",
		District:         "Bairro",
		State:            domains.SANTA_CANTARINA,
		City:             "Florianópolis",
	}
	fmt.Println(business.ToJson())

	productBankAccount := domains.ProductBankAccount{
		Tipo: "CONTA_CORRENTE",
	}
	fmt.Println(productBankAccount.ToJson())

	proposalBankAccount := domains.ProposalBankAccount{
		Mother:             "Julia",
		Gender:             domains.FEMININO,
		Natianality:        domains.BRASILEIRA,
		HomeTownState:      domains.SANTA_CANTARINA,
		RelationshipStatus: domains.CASADO,
		Identity:           domains.CNH,
		Adrres:             address,
		Business:           business,
		Products:           productBankAccount,
	}
	fmt.Println(proposalBankAccount.ToJson())

	assert.Equal(t, "Julia", proposalBankAccount.Mother)
	assert.Equal(t, domains.FEMININO, proposalBankAccount.Gender)
	assert.Equal(t, domains.BRASILEIRA, proposalBankAccount.Natianality)
	assert.Equal(t, domains.SANTA_CANTARINA, proposalBankAccount.HomeTownState)
	assert.Equal(t, domains.CASADO, proposalBankAccount.RelationshipStatus)
	assert.Equal(t, domains.CNH, proposalBankAccount.Identity)
	assert.Equal(t, address, proposalBankAccount.Adrres)
	assert.Equal(t, business, proposalBankAccount.Business)
	assert.Equal(t, productBankAccount, proposalBankAccount.Products)

}
