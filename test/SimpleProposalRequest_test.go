package test

import (
    "modulo/src/requests"
    "modulo/src/domains"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSimpleProposalRequest(t *testing.T){

       address := domains.Address{
       		ZipCode:    "12345678",
       		Address:    "Rua dos Bobos",
       		Number:     "0",
       		Complement: "Casa",
       		District:   "Bairro",
       		State:      domains.SANTA_CANTARINA,
       		City:       "Florianópolis",
       	}

       	business := domains.Business{
        		Occupation:      domains.ASSALARIADO,
        		Profession:      domains.ACUMPURISTA,
        		CompanyName:      "Empresa",
        		Income:           "1000",
        		EmploymentSince: domains.LESS_THAN_SIX_MONTHS,
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

       productBankAccount := domains.ProductBankAccount{
       		Tipo: "CONTA_CORRENTE",
       	}

        proposalBankAccount := domains.ProposalBankAccount{
        		Mother:             "Julia",
        		Gender:             domains.FEMININO,
        		Nationality:        domains.BRASILEIRA,
        		HomeTownState:      domains.GOIAS,
        		RelationshipStatus: domains.CASADO,
        	    Address:            address,
        		Business:           business,
        		Products:           productBankAccount,
        }

        assert.Equal(t, "Julia", proposalBankAccount.Mother)
        assert.Equal(t, domains.FEMININO, proposalBankAccount.Gender)
        assert.Equal(t, domains.BRASILEIRA, proposalBankAccount.Nationality)
        assert.Equal(t, domains.GOIAS, proposalBankAccount.HomeTownState)
        assert.Equal(t, domains.CASADO, proposalBankAccount.RelationshipStatus)

        assert.Equal(t, address, proposalBankAccount.Address)
        assert.Equal(t, business, proposalBankAccount.Business)
        assert.Equal(t, productBankAccount, proposalBankAccount.Products)


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