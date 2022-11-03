package domains

type Document struct {
	tipo DocumentType
	mimeType MimeType
	name string
	base64 string
}