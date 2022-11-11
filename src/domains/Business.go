package domains

import "encoding/json"

type Business struct {
	Occupation_      Occupation
	Profession_      Profession
	CompanyName      string
	Income           string
	EmploymentSince_ EmploymentSince
	Payday           string
	BenefitNumber    string
	ZipCode          string
	Address          string
	Number           string
	Complement       string
	District         string
	State            HomeTownState
	City             string
}

func (a *Business) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
