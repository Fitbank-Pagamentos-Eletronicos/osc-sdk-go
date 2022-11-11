package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelineMatchLoan(t *testing.T) {

	matchLoan := domains.MatchLoan{
		ProductId:      1,
		Name:           "Teste",
		Logo:           "Teste",
		MinValue:       1,
		MaxValue:       1000,
		MinInstallment: 1,
		MaxInstallment: 12,
		MonthlyTax:     1.0,
	}

	fmt.Println(matchLoan.ToJson())

	matchCard := domains.MatchCard{
		ProductId: 2,
		Name:      "Teste",
		Logo:      "Teste",
		Annuity:   1.0,
		Network:   domains.VISA,
	}
	fmt.Println(matchCard.ToJson())

	matchAuto := domains.MatchAuto{
		ProductId:      3,
		Name:           "Teste",
		Logo:           "Teste",
		MinValue:       1,
		MaxValue:       1000,
		MinInstallment: 1,
		MaxInstallment: 12,
		MonthlyTax:     2.0,
	}
	fmt.Println(matchAuto.ToJson())

	matchHome := domains.MatchHome{
		ProductId:      5,
		Name:           "Teste",
		Logo:           "Teste",
		MinValue:       1,
		MaxValue:       1000,
		MinInstallment: 1,
		MaxInstallment: 16,
		MonthlyTax:     1.0,
	}
	fmt.Println(matchHome.ToJson())

	matchs := domains.Match{
		MatchLoan: matchLoan,
		MatchCard: matchCard,
		MatchAuto: matchAuto,
		MatchHome: matchHome,
	}

	pipelineMatchLoan := domains.PipelineMatchLoan{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "12345678901",
		Name:        "Emanuel da Silva",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
		Matches:     matchs,
	}
	fmt.Println(pipelineMatchLoan.ToJson())

	assert.Equal(t, "1", pipelineMatchLoan.Id)
	assert.Equal(t, "Aprovado", pipelineMatchLoan.Status)
	assert.Equal(t, "12345678901", pipelineMatchLoan.Cpf)
	assert.Equal(t, "Emanuel da Silva", pipelineMatchLoan.Name)
	assert.Equal(t, "2020-01-01", pipelineMatchLoan.DateCreated)
	assert.Equal(t, "2020-01-01", pipelineMatchLoan.LastUpdated)
	assert.Equal(t, matchs, pipelineMatchLoan.Matches)
}
