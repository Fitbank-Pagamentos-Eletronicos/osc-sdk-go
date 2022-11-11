package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusiness(t *testing.T) {
	business := domains.Business{
		Occupation_:      domains.AUTONOMO,
		Profession_:      domains.ACOUGUEIRO,
		CompanyName:      "Empresa de Teste",
		Income:           "1000",
		EmploymentSince_: domains.LESS_THAN_SIX_MONTHS,
		Payday:           "10",
		BenefitNumber:    "123456789",
		ZipCode:          "12345678",
		Address:          "Rua de Teste",
		Number:           "123",
		Complement:       "Casa",
		District:         "Bairro de Teste",
		State:            domains.ALAGOAS,
		City:             "Maceió",
	}

	fmt.Println(business.ToJson())

	assert.Equal(t, domains.AUTONOMO, business.Occupation_)
	assert.Equal(t, domains.ACOUGUEIRO, business.Profession_)
	assert.Equal(t, "Empresa de Teste", business.CompanyName)
	assert.Equal(t, "1000", business.Income)
	assert.Equal(t, domains.LESS_THAN_SIX_MONTHS, business.EmploymentSince_)
	assert.Equal(t, "10", business.Payday)
	assert.Equal(t, "123456789", business.BenefitNumber)
	assert.Equal(t, "12345678", business.ZipCode)
	assert.Equal(t, "Rua de Teste", business.Address)
	assert.Equal(t, "123", business.Number)
	assert.Equal(t, "Casa", business.Complement)
	assert.Equal(t, "Bairro de Teste", business.District)
	assert.Equal(t, domains.ALAGOAS, business.State)
	assert.Equal(t, "Maceió", business.City)
}
