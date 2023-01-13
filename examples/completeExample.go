package examples

import (
	"fmt"
	"modulo/src/domains"
	"modulo/src/requests"
)

func main() {
	var instance, _ = requests.CreateInstance("", "", "dafault")

	instance.SetResponseListening(func(pipeline domains.Pipeline, err bool) {

		switch pipeline.Status {
		case domains.SIGNUP_ANALISIS:
			fmt.Printf("Async %s cadastro em analise", pipeline.Id)
		case domains.SIGNUP_COMPLETED:
			fmt.Printf("Async %s cadastro em completo", pipeline.Id)
			// Envia chamada da proposta
			proposal(pipeline.Id)
		case domains.SIGNUP_DENIED:
			fmt.Printf("Async %s cadastro regeitado", pipeline.Id)
		case domains.PROPOSAL_ANALISIS:
			fmt.Printf("Async %s proposta em analise", pipeline.Id)
		case domains.PROPOSAL_CREATED:
			fmt.Printf("Async %s proposta em completo", pipeline.Id)
			for _, proposal := range pipeline.Proposals {
				fmt.Println(proposal)
			}

			// TODO
			//proposal.PendentDocuments
			//sendDocuments
			//
			//proposal.HasContracts
			//getContracts
			//signContractsVis

		case domains.PROPOSAL_DENIED:
			fmt.Printf("Async %s proposta regeitado", pipeline.Id)
		}
	})

	signup()
}

func signup() {
	var data = domains.SimpleSignup{
		Cpf:            "720.825.560-18",
		Name:           "Carlos Henrique",
		Birthday:       "2002-01-19",
		Email:          "CH624323242@email.com",
		Phone:          "63980030021",
		ZipCode:        "74180040",
		HasCreditCart:  true,
		HasRestriction: true,
		HasOwnHouse:    true,
		HasVehicle:     true,
		HasAndroid:     true,
		LogData: domains.LogData{
			Latitude:       -16.678,
			Longitude:      -49.256,
			OccurrenceDate: "2019-08-21T14:31:17.459Z",
			UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36",
			Ip:             "0.0.0.0",
			Mac:            "00:00:00:00:00:00",
		},
	}

	var instance, _ = requests.GetInstance("dafault")
	var pipeline = instance.SimpleSignup(data)

	fmt.Printf("%s", pipeline.Id)
}

func proposal(pipelineId string) {

	var data = domains.ProposalReq{
		Mother:             "Fulana Mãe",
		Gender:             domains.FEMININO,
		Nationality:        domains.BRASILEIRO,
		HomeTownState:      domains.GOIAS,
		HomeTownCity:       "Goiânia",
		Education:          domains.POS_GRADUACAO,
		RelationshipStatus: domains.CASADO,
		PhoneLandLine:      "6232345678",

		Identity: domains.Identity{
			Type:        domains.RG,
			Number:      "123456",
			Issuer:      domains.SSP,
			State:       domains.GOIAS,
			IssuingDate: "2010-01-01",
		},

		Address: domains.Address{
			ZipCode:    "74000000",
			Address:    "Cidade de Goiânia",
			Number:     "0",
			Complement: "toda a cidade",
			District:   "geral",
			State:      domains.GOIAS,
			City:       "Goiânia",
			HomeType:   domains.ALUGADA,
			HomeSince:  domains.MAIOR_2_ANOS,
		},

		Vehicle: domains.Vehicle{
			LicensePlate: "XXX0000",
		},

		ConsumerUnit: domains.ConsumerUnit{
			Number: "000000000",
		},

		Business: domains.Business{
			Occupation:      domains.ASSALARIADO,
			Profession:      domains.ADMINISTRADOR,
			CompanyName:     "Empresa 1",
			Phone:           "6239413142",
			Income:          "1500",
			EmploymentSince: domains.LESS_THAN_SIX_MONTHS,
			Payday:          "10",
			BenefitNumber:   "123456789",
			ZipCode:         "74180040",
			Address:         "Cidade de Goiânia",
			Number:          "123",
			Complement:      "Casa 1",
			District:        "Setor Central",
			State:           domains.GOIAS,
			City:            "Goiânia",
		},

		Bank: domains.Bank{
			Bank:    domains.BANCO_DO_BRASIL,
			Tipo:    domains.CONTA_CORRENTE_INDIVIDUAL,
			Agency:  "0001",
			Account: "123456789",
		},

		Reference: domains.Reference{
			Name:  "Joana Maria",
			Phone: "62987832321",
		},

		Products: domains.Products{
			Type:               domains.REFINANCING_HOME,
			Value:              15000.00,
			Installments:       12,
			Network:            domains.MASTERCARD,
			Payday:             "10",
			VehicleBrand:       "Volkswagen",
			VehicleModel:       "Gol",
			VehicleModelYear:   "2010",
			RealEstateType:     domains.HOUSE,
			RealEstateValue:    148000.00,
			OutstandingBalance: 50000.00,
		},
	}

	var instance, _ = requests.GetInstance("dafault")
	var pipeline = instance.Proposal(pipelineId, data)

	fmt.Printf("%s", pipeline.Id)
}
