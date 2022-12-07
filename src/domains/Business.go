package domains

import "encoding/json"

type Business struct {
	Occupation      Occupation
	Profession      Profession
	CompanyName      string `json:"companyName"`
	Phone            string `json:"phone"`
	Income           string `json:"income"`
	EmploymentSince  EmploymentSince
	Payday           string `json:"payday"`
	BenefitNumber    string `json:"benefitNumber"`
	ZipCode          string `json:"zipCode"`
	Address          string `json:"address"`
	Number           string `json:"number"`
	Complement       string `json:"complement"`
	District         string `json:"district"`
	State            HomeTownState
	City             string `json:"city"`
}

func (a *Business) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
