package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuto(t *testing.T) {
	auto := domains.Auto{
		CustomerServiceNumber: "123",
		Tipo:                  domains.LOAN,
		Product:               "CARRO",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:      domains.IDENTITY_FRONT_,
		DateCreated:           "2019-08-21T14:31:17.459Z",
		LastUpdated:           "2019-08-21T14:31:17.459Z",
		Value:                 1000.00,
		Installments:          123,
		MonthlyTax:            100.00,
		InstallmentsValue:     100.00,
		IofValeu:              100.00,
		GrossValeu:            100.00,
		FirstPaymentDate:      "2019-08-21T14:31:17.459Z",
		Cet:                   100.00,
		ReleaseDate:           "2019-08-21T14:31:17.459Z",
	}
	
	fmt.Println(auto.ToJson())

	assert.Equal(t, "123", auto.CustomerServiceNumber)
	assert.Equal(t, domains.LOAN, auto.Tipo)
	assert.Equal(t, "CARRO", auto.Product)
	assert.Equal(t, 1, auto.ProductId)
	assert.Equal(t, true, auto.HasDocuments)
	assert.Equal(t, true, auto.HasContracts)
	assert.Equal(t, "logo", auto.Logo)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, auto.LastStatus)
	assert.Equal(t, domains.IDENTITY_FRONT_, auto.PendentDocuments)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", auto.DateCreated)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", auto.LastUpdated)
	assert.Equal(t, 1000.00, auto.Value)
	assert.Equal(t, 123, auto.Installments)
	assert.Equal(t, 100.00, auto.MonthlyTax)
	assert.Equal(t, 100.00, auto.InstallmentsValue)
	assert.Equal(t, 100.00, auto.IofValeu)
	assert.Equal(t, 100.00, auto.GrossValeu)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", auto.FirstPaymentDate)
	assert.Equal(t, 100.00, auto.Cet)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", auto.ReleaseDate)

}
