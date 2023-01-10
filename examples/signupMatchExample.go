package examples

import (
	"fmt"
	"modulo/src/domains"
	"modulo/src/osc"
)

func main() {
	var instance, _ = osc.CreateInstance("", "", "dafault")
	var data = domains.SignupMatch{
		Cpf:            "720.825.560-18",
		Name:           "Carlos Henrique",
		Birthday:       "2002-01-19",
		Email:          "CH624323242@email.com",
		Phone:          "63980030021",
		ZipCode:        "74180040",
		Education:      domains.POS_GRADUACAO,
		Banks:          domains.BANCO_DO_BRASIL,
		Occupation:     domains.ASSALARIADO,
		Income:         1045,
		HasCreditCart:  true,
		HasRestriction: true,
		HasOwnHouse:    true,
		HasVehicle:     true,
		HasAndroid:     true,

		Products: domains.Products{
			Type:               domains.LOAN,
			Value:              7000.00,
			Installments:       12,
			Network:            domains.MASTERCARD,
			Payday:             "1",
			VehicleBrand:       "Volkswagen",
			VehicleModel:       "Gol",
			VehicleModelYear:   "2010",
			CodeFipe:           "123456789",
			VehicleFipeValue:   10000.00,
			RealEstateType:     domains.HOUSE,
			RealEstateValue:    148000.00,
			OutstandingBalance: 50000.00,
		},

		LogData: domains.LogData{
			Latitude:       -16.678,
			Longitude:      -49.256,
			OccurrenceDate: "2019-08-21T14:31:17.459Z",
			UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36",
			Ip:             "0.0.0.0",
			Mac:            "00:00:00:00:00:00",
		},
	}

	var pipeline = instance.SignupMatch(data)

	fmt.Printf("%s", pipeline.Id)
}
