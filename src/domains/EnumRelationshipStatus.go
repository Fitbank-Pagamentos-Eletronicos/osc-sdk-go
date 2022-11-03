package domains

type RelationshipStatus struct {
	relationshipStatus string
}

const (
	CASADO = RelationshipStatus{"CASADO"}
	DIVORCIADO = RelationshipStatus{"DIVORCIADO"}
	VIUVO = RelationshipStatus{"VIUVO"}
	SEPARADO = RelationshipStatus{"SEPARADO"}
	SOLTEIRO = RelationshipStatus{"SOLTEIRO"}
	COMPANHEIRO = RelationshipStatus{"COMPANHEIRO"}
	UNIAO_ESTAVEL = RelationshipStatus{"UNIAO_ESTAVEL"}
	SEPARADA_JUDICIALMENTE = RelationshipStatus{"SEPARADA_JUDICIALMENTE"}

)