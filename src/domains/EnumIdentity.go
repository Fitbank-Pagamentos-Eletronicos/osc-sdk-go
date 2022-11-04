package domains

type IdentityType struct {
	IdentityType  string
}

const (
	RG = Identity{"RG"}
	CNH = Identity{"CNH"}
	PASSAPORTE = Identity{"PASSAPORTE"}
	CARTEIRA_CRC = Identity{"CARTEIRA CRC"}
	CARTEIRA_CREA = Identity{"CARTEIRA CREA"}
	CARTEIRA_OAB = Identity{"CARTEIRA OAB"}
	CARTEIRA_CRE = Identity{"CARTEIRA CRE"}

)