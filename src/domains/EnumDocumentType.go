package domains

type DocumentType struct{
	documentType string
}

const (
	SELF = DocumentType{"SELF"}
	IDENTITY_FRONT = DocumentType{"IDENTITY_FRONT"}
	IDENTITY_BACK = DocumentType{"IDENTITY_BACK"}
	ADDRESS_PROOF = DocumentType{"ADDRESS_PROOF"}
	INCOME_PROOF = DocumentType{"INCOME_PROOF"}

)




