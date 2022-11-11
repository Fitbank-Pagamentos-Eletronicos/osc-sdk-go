package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredit(t *testing.T) {
	credit := domains.Credit{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.REFINANCING_HOME,
		Product:               "Crédito Imobiliário",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:      domains.IDENTITY_FRONT_,
		DateCreated:           "01/01/2020",
		LastUpdated:           "01/01/2020",
	}

	fmt.Println(credit.ToJson())

	assert.Equal(t, "123456789", credit.CustomerServiceNumber)
	assert.Equal(t, domains.REFINANCING_HOME, credit.Tipo)
	assert.Equal(t, "Crédito Imobiliário", credit.Product)
	assert.Equal(t, 1, credit.ProductId)
	assert.Equal(t, true, credit.HasDocuments)
	assert.Equal(t, true, credit.HasContracts)
	assert.Equal(t, "logo", credit.Logo)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, credit.LastStatus)
	assert.Equal(t, domains.IDENTITY_FRONT_, credit.PendentDocuments)
	assert.Equal(t, "01/01/2020", credit.DateCreated)
	assert.Equal(t, "01/01/2020", credit.LastUpdated)

}
