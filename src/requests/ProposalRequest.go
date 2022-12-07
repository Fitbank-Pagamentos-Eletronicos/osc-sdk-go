package requests

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
    "modulo/src/domains"
)

var dataProposal = domains.Proposal {
    Mother:     "Fulana Mãe",
    Gender:    domains.FEMININO,
    Nationality: domains.BRASILEIRA,
    HomeTownState: domains.GOIAS,
    HomeTownCity: "Goiânia",
    Education: domains.POS_GRADUACAO,
    RelationshipStatus: domains.CASADO,
    PhoneLandLine: "62987832321",

    Identity: domains.Identity{
        Tipo: domains.RG,
        Number: "123456789",
        Issuer: domains.SSP,
        State: domains.GOIAS,
        IssuingDate: "2010-01-01",

    },

    Address: domains.Address{
        ZipCode: "74180040",
        Address: "Rua 1",
        Number: "123",
        Complement: "Casa 1",
        District: "Setor Central",
        State: domains.GOIAS,
        City: "Goiânia",
        HomeType: domains.ALUGADA,
        HomeSince: domains.MAIOR_2_ANOS,
    },

    Vehicle: domains.Vehicle{
        LicensePlate: "ABC1234",
    },

    ConsumerUnit: domains.ConsumerUnit{
        Number: "123456789",
    },

    Business: domains.Business{
        Occupation: domains.ASSALARIADO,
        Profession: domains.ADVOGADO,
        CompanyName: "Empresa 1",
        Phone: "62987832321",
        Income: "1000.00",
        EmploymentSince: domains.LESS_THAN_SIX_MONTHS,
        Payday: "10",
        BenefitNumber: "123456789",
        ZipCode: "74180040",
        Address: "Rua 1",
        Number: "123",
        Complement: "Casa 1",
        District: "Setor Central",
        State: domains.GOIAS,
        City: "Goiânia",
    },

    Bank: domains.Bank{
        Bank: domains.BANCO_DO_BRASIL,
        Tipo: domains.CONTA_CORRENTE_INDIVIDUAL,
        Agency: "0001",
        Account: "123456789",
    },

    Reference: domains.Reference{
        Name: "Joana Maria",
        Phone: "62987832321",
    },

    Products: domains.Products{
        ProductLoan: domains.ProductLoan{
            Tipo: domains.LOAN,
            Value: 7000.00,
            Installments: 12,
        },

        ProductCard: domains.ProductCard{
            Tipo: domains.CARD,
            Network: domains.MASTERCARD,
            Payday: "10",
        },

        ProductAuto: domains.ProductAuto{
          Tipo: domains.REFINANCING_AUTO,
          Value: 8000.00,
          VehicleBrand: "Volkswagen",
          VehicleModel: "Gol",
          Installments: 12,
          VehicleModelYear: "2010",
          CodeFipe: "123456789",
          VehicleFipeValue: 10000.00,
        },

        ProductHome: domains.ProductHome{
            Tipo: domains.REFINANCING_HOME,
            Value: 15000.00,
            Installments: 12,
            RealEstateType_: domains.HOUSE,
            RealEstateValue: 148000.00,
            OutstandingBalance: 50000.00,
        },
    },

}



func ProposalRequest() domains.Proposal{
   url := "https://demo-api.easycredito.com.br/api/external//v2.1/process/proposal/"
   method := "POST"

    simpleProposalToJson, _ := json.Marshal(dataProposal)
    payload := strings.NewReader(string(simpleProposalToJson))


     client := &http.Client {}
     req, err := http.NewRequest(method, url, payload)

     if err != nil {
       fmt.Println(err)
       return domains.Proposal{}
     }
     req.Header.Add("Content-Type", "application/json")
     req.Header.Add("Accept", "application/json")
     req.Header.Add("Authorization", "Bearer " + GetToken())

     res, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
            return domains.Proposal{}
        }

        defer res.Body.Close()

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println(err)
            return domains.Proposal{}
        }
        fmt.Println(string(body))

        var proposal domains.Proposal
        json.Unmarshal(body, &proposal)
        return proposal

}