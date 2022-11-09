package domains

type DocumentResponse struct {
	Tipo     DocumentType
	MimeType MimeType
	Name     string
	Url      string
}
