package domains

import "encoding/json"

type DocumentResponse struct {
	Tipo     DocumentType
	MimeType MimeType
	Name     string
	Url      string
}

func (a *DocumentResponse) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
