package test

import (
	"fmt"
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
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

	reference := domains.Reference{
		Name:  "Referência",
		Phone: "123456789",
	}

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

	assert.Equal(t, proposal.Mother, "Maria")
	assert.Equal(t, proposal.Gender, domains.FEMININO)
	assert.Equal(t, proposal.Nationality, domains.BRASILEIRO)
	assert.Equal(t, proposal.HomeTownState, domains.ACRE)
	assert.Equal(t, proposal.HomeTownCity, "Rio Branco")
	assert.Equal(t, proposal.Education, domains.ENSINO_FUNDAMENTAL_COMPLETO)
	assert.Equal(t, proposal.RelationshipStatus, domains.CASADO)
	assert.Equal(t, proposal.Identity, identity)
	assert.Equal(t, address, proposal.Address)
	assert.Equal(t, vehicle, proposal.Vehicle)
	assert.Equal(t, consumerUnit, proposal.ConsumerUnit)
	assert.Equal(t, business, proposal.Business)
	assert.Equal(t, bank, proposal.Bank)
	assert.Equal(t, reference, proposal.Reference)
	assert.Equal(t, products, proposal.Products)
}
