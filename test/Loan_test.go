package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoan(t *testing.T) {
	loan := domains.Loan{
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

	fmt.Println(loan.ToJson())

	assert.Equal(t, "123456789", loan.CustomerServiceNumber)
	assert.Equal(t, domains.LOAN, loan.Tipo)
	assert.Equal(t, "Empréstimo", loan.Product)
	assert.Equal(t, 1, loan.ProductId)
	assert.Equal(t, true, loan.HasDocuments)
	assert.Equal(t, true, loan.HasContracts)
	assert.Equal(t, "logo", loan.Logo)
	assert.Equal(t, domains.PRE_PROCESSAMENTO, loan.LastStatus)
	assert.Equal(t, domains.ADDRESS_PROOF_, loan.PendentDocuments)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", loan.DateCreated)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", loan.LastUpdated)
	assert.Equal(t, 1000.00, loan.Value)
	assert.Equal(t, 12, loan.Installments)
	assert.Equal(t, 0.00, loan.MonthlyTax)
	assert.Equal(t, 1000.00, loan.InstallmentsValue)
	assert.Equal(t, 0.00, loan.IofValeu)
	assert.Equal(t, 1000.00, loan.GrossValeu)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", loan.FirstPaymentDate)
	assert.Equal(t, 0.00, loan.Cet)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", loan.ReleaseDate)

}
