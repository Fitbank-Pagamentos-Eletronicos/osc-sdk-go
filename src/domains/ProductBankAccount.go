package domains

import "encoding/json"

type ProductBankAccount struct {
  Tipo string `json:"type"`
	ProductId int `json:productId`

}

func (a *ProductBankAccount) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
