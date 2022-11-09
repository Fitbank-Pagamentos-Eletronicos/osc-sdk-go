package domains

type Identity struct {
	Tipo        IdentityType
	Number      string
	Issuer      IdentityTypeIssuer
	State       HomeTownState
	IssuingDate string
}
