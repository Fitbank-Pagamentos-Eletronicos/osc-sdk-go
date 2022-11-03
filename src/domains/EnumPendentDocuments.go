package domains

type PendentDocuments struct {
	pendentDocuments string
}

const (
	SELF_ = PendentDocuments{"SELF"}
	IDENTITY_FRONT_ = PendentDocuments{"IDENTITY_FRONT"}
	IDENTITY_BACK_ = PendentDocuments{"IDENTITY_BACK"}
	ADDRESS_PROOF_ = PendentDocuments{"ADDRESS_PROOF"}
	INCOME_PROOF_ = PendentDocuments{"INCOME_PROOF"}
)


