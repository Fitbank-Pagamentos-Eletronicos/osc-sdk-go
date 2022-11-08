package domains

type IdentityType struct {
	identityType string
}

var (
	RG            IdentityType = IdentityType{"RG"}
	CNH           IdentityType = IdentityType{"CNH"}
	PASSAPORTE    IdentityType = IdentityType{"PASSAPORTE"}
	CARTEIRA_CRC  IdentityType = IdentityType{"CARTEIRA CRC"}
	CARTEIRA_CREA IdentityType = IdentityType{"CARTEIRA CREA"}
	CARTEIRA_OAB  IdentityType = IdentityType{"CARTEIRA OAB"}
	CARTEIRA_CRE  IdentityType = IdentityType{"CARTEIRA CRE"}
)

