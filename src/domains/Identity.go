package domains


type Identity struct {
	tipo  IdentityType
	number string
	issuer IdentityTypeIssuer
	state HomeTownState
	issuingDate string
}