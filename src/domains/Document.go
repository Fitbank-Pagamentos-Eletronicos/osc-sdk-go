package domains

import json2 "encoding/json"

type Document struct {
	Type     DocumentType `json:"type"`
	MimeType MimeType     `json:"mimeType"`
	Name     string       `json:"name"`
	Base64   string       `json:"base64"`
}

func (a *Document) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
