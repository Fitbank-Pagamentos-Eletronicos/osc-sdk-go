package domains

import "encoding/json"

type DetailsOfCET struct {
	InterestPercentage   string `json:"interestPercentage"`
	TaxPercentage        string `json:"taxPercentage"`
	PercentageOfFees     string `json:"percentageOfFees"`
	PercentageOfServices string `json:"percentageOfServices"`
}

func (a *DetailsOfCET) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
