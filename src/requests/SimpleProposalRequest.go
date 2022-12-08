package requests


import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "modulo/src/domains"

)

var dataSimpleProposal = domains.ProposalBankAccount {
    Mother:     "Fulana Mãe",
    Gender: domains.FEMININO,
    Nationality: domains.BRASILEIRA,
    HomeTownState: domains.GOIAS,
    RelationshipStatus: domains.CASADO,

    Address: domains.Address{
        ZipCode: "74000000",
        Address: "Cidade de Goiânia",
        Number: "0",
        Complement: "Casa 1",
        District: "Geral",
        State: domains.GOIAS,
        City: "Goiânia",
    }
    Business: domains.Business{
        Income: "1000.00",
    }

    Products: domains.Products{

    }


}