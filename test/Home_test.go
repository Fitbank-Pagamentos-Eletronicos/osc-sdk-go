package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	home := domains.Home{
		CustomerServiceNumber: "123456789",
		Tipo:                  domains.LOAN,
		Product:               "Empréstimo",
		ProductId:             1,
		HasDocuments:          true,
		HasContracts:          true,
		Logo:                  "logo",
		LastStatus:            domains.PRE_PROCESSAMENTO,
		PendentDocuments:      domains.ADDRESS_PROOF_,
		DateCreated:           "2019-08-21T14:31:17.459Z",
		LastUpdated:           "2019-08-21T14:31:17.459Z",
		Value:                 1000.00,
		Installments:          12,
		MonthlyTax:            0.00,
		InstallmentsValue:     1000.00,
		IofValeu:              0.00,
		GrossValeu:            1000.00,
		FirstPaymentDate:      "2019-08-21T14:31:17.459Z",
		Cet:                   0.00,
		ReleaseDate:           "2019-08-21T14:31:17.459Z",
	}

	fmt.Println(home.ToJson())

	assert.Equal(t, "123456789", home.CustomerServiceNumber)
	assert.Equal(t, domains.LOAN, home.Tipo)
	assert.Equal(t, "Empréstimo", home.Product)
	assert.Equal(t, 1, home.ProductId)
	assert.Equal(t, true, home.HasDocuments)
	assert.Equal(t, true, home.HasContracts)
	assert.Equal(t, "logo", home.Logo)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, home.LastStatus)
	assert.Equal(t, domains.ADDRESS_PROOF_, home.PendentDocuments)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", home.DateCreated)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", home.LastUpdated)
	assert.Equal(t, 1000.00, home.Value)
	assert.Equal(t, 12, home.Installments)
	assert.Equal(t, 0.00, home.MonthlyTax)
	assert.Equal(t, 1000.00, home.InstallmentsValue)
	assert.Equal(t, 0.00, home.IofValeu)
	assert.Equal(t, 1000.00, home.GrossValeu)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", home.FirstPaymentDate)
	assert.Equal(t, 0.00, home.Cet)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", home.ReleaseDate)

}
