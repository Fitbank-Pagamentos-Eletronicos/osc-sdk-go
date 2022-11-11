package domains

import "encoding/json"

type Document struct {
	Tipo      DocumentType
	MimeType_ MimeType
	Name      string
	Base64    string
}

func (a *Document) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
