package domains

import json2 "encoding/json"

type InstallmentsSimulation struct {
	Account                   string                     `json:"account"`
	GrossValue                string                     `json:"grossValue"`
	SecureValue               string                     `json:"secureValue"`
	InsuranceValue            string                     `json:"insuranceValue"`
	CADValue                  string                     `json:"cadValue"`
	IOFValue                  string                     `json:"iofValue"`
	NetValue                  string                     `json:"netValue"`
	InterestRate              string                     `json:"interestRate"`
	CET                       string                     `json:"cet"`
	AnnualCET                 string                     `json:"annualCet"`
	ReleaseDate               string                     `json:"releaseDate"`
	OriginalDueDate           string                     `json:"originalDueDate"`
	FinalDueDate              string                     `json:"finalDueDate"`
	NumberOfInstallments      int                        `json:"numberOfInstallments"`
	Installments              ForecastOfInstallmentsBody `json:"installments"`
	TotalAmortization         string                     `json:"totalAmortization"`
	TotalServiceCharge        string                     `json:"totalServiceCharge"`
	TotalBrokerage            string                     `json:"totalBrokerage"`
	TotalInsurance            string                     `json:"totalInsurance"`
	TotalInterest             string                     `json:"totalInterest"`
	TotalAmountOfInstallments string                     `json:"totalAmountOfInstallments"`
	DetailsCET                DetailsOfCET               `json:"detailsCET"`
}

func (a *InstallmentsSimulation) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
