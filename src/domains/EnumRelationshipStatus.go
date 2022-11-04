package domains

type RelationshipStatus struct {
	relationshipStatus string
}

var (
	CASADO                 RelationshipStatus = RelationshipStatus{"CASADO"}
	DIVORCIADO             RelationshipStatus = RelationshipStatus{"DIVORCIADO"}
	VIUVO                  RelationshipStatus = RelationshipStatus{"VIUVO"}
	SEPARADO               RelationshipStatus = RelationshipStatus{"SEPARADO"}
	SOLTEIRO               RelationshipStatus = RelationshipStatus{"SOLTEIRO"}
	COMPANHEIRO            RelationshipStatus = RelationshipStatus{"COMPANHEIRO"}
	UNIAO_ESTAVEL          RelationshipStatus = RelationshipStatus{"UNIAO_ESTAVEL"}
	SEPARADA_JUDICIALMENTE RelationshipStatus = RelationshipStatus{"SEPARADA_JUDICIALMENTE"}
)
