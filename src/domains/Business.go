package domains

import json2 "encoding/json"

type Business struct {
	Occupation      Occupation      `json:"occupation"`
	Profession      Profession      `json:"profession"`
	CompanyName     string          `json:"companyName"`
	Phone           string          `json:"phone"`
	Income          string          `json:"income"`
	EmploymentSince EmploymentSince `json:"employmentSince"`
	Payday          string          `json:"payday"`
	BenefitNumber   string          `json:"benefitNumber"`
	ZipCode         string          `json:"zipCode"`
	Address         string          `json:"address"`
	Number          string          `json:"number"`
	Complement      string          `json:"complement"`
	District        string          `json:"district"`
	State           HomeTownState   `json:"state"`
	City            string          `json:"city"`
}

func (a *Business) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
