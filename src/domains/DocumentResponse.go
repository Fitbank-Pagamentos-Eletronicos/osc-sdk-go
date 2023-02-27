package domains

import json2 "encoding/json"

type DocumentResponse struct {
	Type     DocumentType `json:"type"`
	MimeType MimeType     `json:"mimeType"`
	Name     string       `json:"name"`
	Url      string       `json:"url"`
}

func (a *DocumentResponse) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
