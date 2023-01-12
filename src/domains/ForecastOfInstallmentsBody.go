package domains

import json2 "encoding/json"

type ForecastOfInstallmentsBody struct {
	ParcelNumber                string `json:"parcelNumber"`
	ExpirationDate              string `json:"expirationDate"`
	AmortizationAmount          string `json:"AmortizationAmount"`
	CorrectionAmortizationValue string `json:"correctionAmortizationValue"`
	InterestValue               string `json:"interestValue"`
	InsuranceValue              string `json:"insuranceValue"`
	BankRate                    string `json:"bankRate"`
	InstallmentAmount           string `json:"installmentAmount"`
	PreviousBalanceValue        string `json:"previousBalanceValue"`
	CapitalizedInterestValue    string `json:"capitalizedInterestValue"`
	CurrentBalanceValue         string `json:"currentBalanceValue"`
}

func (a *ForecastOfInstallmentsBody) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
