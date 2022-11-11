package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSignup(t *testing.T) {
	logData := domains.LogData{
		Latitude:      12.3,
		Longitude:     45.7,
		OccurenceData: "2019-01-01T00:00:00Z",
		UserAgent:     "Mozilla/5.0",
		Ip:            "1.0.0.0",
		Mac:           "00:00:00:00:00:00",
	}

	fmt.Println(logData.ToJson())

	simpleSignup := domains.SimpleSignup{
		Cpf:            "123456789",
		Name:           "Van",
		Birthday:       "01/01/2000",
		Email:          "van@gmail.com",
		Phone:          "123456789",
		ZipCode:        "12345678",
		HasCreditCart:  true,
		HasRestriction: true,
		HasOwnHouse:    true,
		HasVehicle:     true,
		HasAndroid:     true,
		LogaData:       logData,
	}

	fmt.Println(simpleSignup.ToJson())

	assert.Equal(t, "123456789", simpleSignup.Cpf)
	assert.Equal(t, "Van", simpleSignup.Name)
	assert.Equal(t, "01/01/2000", simpleSignup.Birthday)
	assert.Equal(t, "van@gmail.com", simpleSignup.Email)
	assert.Equal(t, "123456789", simpleSignup.Phone)
	assert.Equal(t, "12345678", simpleSignup.ZipCode)
	assert.Equal(t, true, simpleSignup.HasCreditCart)
	assert.Equal(t, true, simpleSignup.HasRestriction)
	assert.Equal(t, true, simpleSignup.HasOwnHouse)
	assert.Equal(t, true, simpleSignup.HasVehicle)
	assert.Equal(t, true, simpleSignup.HasAndroid)
	assert.Equal(t, logData, simpleSignup.LogaData)
}
