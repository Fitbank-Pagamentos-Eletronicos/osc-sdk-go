package domains

type MimeType struct {
	mimeType string
}

var (
	IMAGE_JPEG      MimeType = MimeType{"image/jpeg"}
	IMAGE_PNG       MimeType = MimeType{"image/png"}
	APPLICATION_PDF MimeType = MimeType{"application/pdf"}
)
