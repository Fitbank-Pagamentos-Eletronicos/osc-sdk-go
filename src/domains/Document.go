package domains

import "encoding/json"

type Document struct {
	Tipo      DocumentType `json: "tipo"`
	MimeType  MimeType `json: "mimeType"`
	Name      string `json: "name"`
	Base64    string `json: base64`
}

func (a *Document) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
