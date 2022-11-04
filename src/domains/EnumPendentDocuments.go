package domains

type PendentDocuments struct {
	pendentDocuments string
}

var (
	SELF_           PendentDocuments = PendentDocuments{"SELF"}
	IDENTITY_FRONT_ PendentDocuments = PendentDocuments{"IDENTITY_FRONT"}
	IDENTITY_BACK_  PendentDocuments = PendentDocuments{"IDENTITY_BACK"}
	ADDRESS_PROOF_  PendentDocuments = PendentDocuments{"ADDRESS_PROOF"}
	INCOME_PROOF_   PendentDocuments = PendentDocuments{"INCOME_PROOF"}
)
