package domains

type MimeType struct{
	mimeType string
}

const (
	IMAGE_JPEG = MimeType{"image/jpeg"}
	IMAGE_PNG = MimeType{"image/png"}
	APPLICATION_PDF = MimeType{"application/pdf"}
)