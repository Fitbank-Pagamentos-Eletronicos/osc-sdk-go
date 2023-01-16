package domains

import json2 "encoding/json"

type ProductWorkingCapital struct {
	Type                         ProductType                  `json:"type"`
	Installments                 int                          `json:"installments"`
	Value                        float64                      `json:"value"`
	CNPJ                         string                       `json:"cnpj"`
	BusinessProfession           BusinessProfession           `json:"businessProfession"`
	EmployeesCount               EmployeesCount               `json:"employeeAccount"`
	BusinessIncomeCnpj           string                       `json:"businessIncomeCnpj"`
	WorkingCapitalLoanObjectives WorkingCapitalLoanObjectives `json:"workingCapitalLoanObjectives"`
	Bank                         Bank                         `json:"bank"`
	AccountType                  Bank                         `json:"accountType"`
	Agency                       Bank                         `json:"agency"`
	Account                      Bank                         `json:"account"`
}

func (a *ProductWorkingCapital) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
