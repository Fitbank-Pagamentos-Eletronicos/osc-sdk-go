package domains

type EducationLevel struct {
	nome string
}

//Enum de Nivel de Escolaridade
const (
	NAO_ALFABETIZADO = EducationLevel{"Não alfabetizado"}
	ENSINO_FUNDAMENTAL_INCOMPLETO = EducationLevel{"Ensino Fundamental incompleto"}
	ENSINO_FUNDAMENTAL_COMPLETO = EducationLevel{"Ensino Fundamental completo"}
	ENSINO_MEDIO_INCOMPLETO = EducationLevel{"Ensino Médio incompleto"}
	ENSINO_MEDIO_COMPLETO = EducationLevel{"Ensino Médio completo"}
	ENSINO_SUPERIOR_INCOMPLETO = EducationLevel{"Ensino Superior incompleto"}
	ENSINO_SUPERIOR_COMPLETO = EducationLevel{"Ensino Superior completo"}
	POS_GRADUACAO = EducationLevel{"Pós-graduação"}
)