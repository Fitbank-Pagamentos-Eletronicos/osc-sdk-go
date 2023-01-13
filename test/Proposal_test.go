package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposal(t *testing.T) {
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

	vehicle := domains.Vehicle{
		LicensePlate: "ABC1234",
	}

	fmt.Println(vehicle.ToJson())

	consumerUnit := domains.ConsumerUnit{
		Number: "123456789",
	}
	fmt.Println(consumerUnit.ToJson())
	business := domains.Business{
		Occupation:      domains.APOSENTADO,
		Profession:      domains.ACUMPURISTA,
		CompanyName:     "Empresa",
		Income:          "1000",
		EmploymentSince: domains.LESS_THAN_SIX_MONTHS,
		Payday:          "10",
		BenefitNumber:   "123456789",
		ZipCode:         "12345678",
		Address:         "Rua dos Bobos",
		Number:          "0",
		Complement:      "Casa",
		District:        "Bairro",
		State:           domains.SANTA_CANTARINA,
		City:            "Florianópolis",
	}
	fmt.Println(business.ToJson())

	bank := domains.Bank{
		Bank:    domains.NU_PAGAMENTOS_S,
		Type:    domains.CONTA_CORRENTE_CONJUNTA,
		Agency:  "1234",
		Account: "123456",
	}
	fmt.Println(bank.ToJson())
	reference := domains.Reference{
		Name:  "Referência",
		Phone: "123456789",
	}

	fmt.Println(reference.ToJson())

	productLoan := domains.ProductLoan{
		Type:         domains.LOAN,
		Value:        1000,
		Installments: 12,
	}
	fmt.Println(productLoan.ToJson())

	productCard := domains.ProductCard{
		Type:    domains.CARD,
		Network: domains.VISA,
		Payday:  "10",
	}
	fmt.Println(productCard.ToJson())

	productAuto := domains.ProductAuto{
		Type:             domains.LOAN,
		Value:            1000.00,
		VehicleBrand:     "Fiat",
		VehicleModel:     "Uno",
		Installments:     12,
		CodeFipe:         "123456",
		VehicleFipeValue: 1000.00,
	}
	fmt.Println(productAuto.ToJson())

	productHome := domains.ProductHome{
		Type:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		RealEstateType:     domains.HOUSE,
		RealEstateValue:    1000.00,
		OutstandingBalance: 1000.00,
	}
	fmt.Println(productHome.ToJson())

	products := domains.Products{
		Type:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		Network:            domains.VISA,
		Payday:             "20",
		VehicleBrand:       "Fiat",
		VehicleModel:       "Uno",
		CodeFipe:           "123",
		VehicleFipeValue:   10000.00,
		RealEstateType:     domains.HOUSE,
		RealEstateValue:    10000.00,
		OutstandingBalance: 1000.00,
	}

	identity := domains.Identity{
		Type:        domains.CNH,
		Number:      "123456789",
		Issuer:      domains.SSP,
		State:       domains.ACRE,
		IssuingDate: "2020-01-01",
	}

	proposal := domains.ProposalReq{
		Mother:             "Maria",
		Gender:             domains.FEMININO,
		Nationality:        domains.BRASILEIRO,
		HomeTownState:      domains.ACRE,
		HomeTownCity:       "Rio Branco",
		Education:          domains.ENSINO_FUNDAMENTAL_COMPLETO,
		RelationshipStatus: domains.CASADO,
		Identity:           identity,
		Address:            address,
		Vehicle:            vehicle,
		ConsumerUnit:       consumerUnit,
		Business:           business,
		Bank:               bank,
		Reference:          reference,
		Products:           products,
	}

	fmt.Println(proposal.ToJson())

	assert.Equal(t, proposal.Mother, "Maria")
	assert.Equal(t, proposal.Gender, domains.FEMININO)
	assert.Equal(t, proposal.Nationality, domains.BRASILEIRO)
	assert.Equal(t, proposal.HomeTownState, domains.ACRE)
	assert.Equal(t, proposal.HomeTownCity, "Rio Branco")
	assert.Equal(t, proposal.Education, domains.ENSINO_FUNDAMENTAL_COMPLETO)
	assert.Equal(t, proposal.RelationshipStatus, domains.CASADO)
	assert.Equal(t, proposal.Identity, domains.CNH)
	assert.Equal(t, address, proposal.Address)
	assert.Equal(t, vehicle, proposal.Vehicle)
	assert.Equal(t, consumerUnit, proposal.ConsumerUnit)
	assert.Equal(t, business, proposal.Business)
	assert.Equal(t, bank, proposal.Bank)
	assert.Equal(t, reference, proposal.Reference)
	assert.Equal(t, products, proposal.Products)
}
