package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupMatch(t *testing.T) {

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

	logData := domains.LogData{
		Latitude:       1.0,
		Longitude:      1.0,
		OccurrenceDate: "2020-01-01",
		UserAgent:      "Chrome",
		Ip:             "0.0.0.0",
		Mac:            "00:00:00:00:00:00",
	}

	signupMatch := domains.SignupMatch{
		Cpf:            "123456789",
		Name:           "John Doe",
		Birthday:       "2020-01-01",
		Phone:          "123456789",
		ZipCode:        "12345",
		Education:      domains.ENSINO_SUPERIOR_COMPLETO,
		Banks:          domains.BANCO_PAN,
		Occupation:     domains.AUTONOMO,
		Income:         1000,
		HasCreditCart:  true,
		HasRestriction: true,
		HasOwnHouse:    true,
		HasVehicle:     true,
		HasAndroid:     true,
		Products:       products,
		LogData:        logData,
	}

	assert.Equal(t, signupMatch.Cpf, "123456789")
	assert.Equal(t, signupMatch.Name, "John Doe")
	assert.Equal(t, signupMatch.Birthday, "2020-01-01")
	assert.Equal(t, signupMatch.Phone, "123456789")
	assert.Equal(t, signupMatch.ZipCode, "12345")
	assert.Equal(t, signupMatch.Education, domains.ENSINO_SUPERIOR_COMPLETO)
	assert.Equal(t, signupMatch.Banks, domains.BANCO_PAN)
	assert.Equal(t, signupMatch.Occupation, domains.AUTONOMO)
	assert.Equal(t, signupMatch.Income, 1000)
	assert.Equal(t, signupMatch.HasCreditCart, true)
	assert.Equal(t, signupMatch.HasRestriction, true)
	assert.Equal(t, signupMatch.HasOwnHouse, true)
	assert.Equal(t, signupMatch.HasVehicle, true)
	assert.Equal(t, signupMatch.HasAndroid, true)
	assert.Equal(t, products, signupMatch.Products)
	assert.Equal(t, logData, signupMatch.LogData)

}
