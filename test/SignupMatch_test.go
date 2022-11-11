package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupMatch(t *testing.T) {
	productLoan := domains.ProductLoan{
		Tipo:         domains.LOAN,
		Value:        1000,
		Installments: 12,
	}

	fmt.Println(productLoan.ToJson())

	productCard := domains.ProductCard{
		Tipo:    domains.CARD,
		Network: domains.VISA,
		Payday:  "10",
	}
	fmt.Println(productCard.ToJson())

	productAuto := domains.ProductAuto{
		Tipo:             domains.LOAN,
		Value:            1000.00,
		VehicleBrand:     "Fiat",
		VehicleModel:     "Uno",
		Installments:     12,
		CodeFipe:         "123456",
		VehicleFipeValue: 1000.00,
	}
	fmt.Println(productAuto.ToJson())

	productHome := domains.ProductHome{
		Tipo:               domains.LOAN,
		Value:              1000.00,
		Installments:       12,
		RealEstateType_:    domains.HOUSE,
		RealEstateValue:    1000.00,
		OutstandingBalance: 1000.00,
	}

	fmt.Println(productHome.ToJson())

	products := domains.Products{
		ProductLoan: productLoan,
		ProductCard: productCard,
		ProductAuto: productAuto,
		ProductHome: productHome,
	}

	fmt.Println(products.ToJson())
	logData := domains.LogData{
		Latitude:      1.0,
		Longitude:     1.0,
		OccurenceData: "2020-01-01",
		UserAgent:     "Chrome",
		Ip:            "0.0.0.0",
		Mac:           "00:00:00:00:00:00",
	}
	fmt.Println(logData.ToJson())

	signupMatch := domains.SignupMatch{
		Cpf:            "123456789",
		Name:           "John Doe",
		Birthday:       "2020-01-01",
		Phone:          "123456789",
		ZipCode:        "12345",
		EducationLevel: domains.ENSINO_SUPERIOR_COMPLETO,
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
	fmt.Println(signupMatch.ToJson())

	assert.Equal(t, signupMatch.Cpf, "123456789")
	assert.Equal(t, signupMatch.Name, "John Doe")
	assert.Equal(t, signupMatch.Birthday, "2020-01-01")
	assert.Equal(t, signupMatch.Phone, "123456789")
	assert.Equal(t, signupMatch.ZipCode, "12345")
	assert.Equal(t, signupMatch.EducationLevel, domains.ENSINO_SUPERIOR_COMPLETO)
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
