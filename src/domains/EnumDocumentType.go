package domains

type DocumentType struct{
	documentType string
}

var (
	SELF DocumentType = DocumentType{"SELF"}
	IDENTITY_FRONT DocumentType = DocumentType{"IDENTITY_FRONT"}
	IDENTITY_BACK DocumentType = DocumentType{"IDENTITY_BACK"}
	ADDRESS_PROOF DocumentType = DocumentType{"ADDRESS_PROOF"}
	INCOME_PROOF DocumentType = DocumentType{"INCOME_PROOF"}

)




