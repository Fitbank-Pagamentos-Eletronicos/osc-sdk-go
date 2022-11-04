package domains

type EducationLevel struct {
	nome string
}

// Enum de Nivel de Escolaridade
var (
	NAO_ALFABETIZADO              EducationLevel = EducationLevel{"Não alfabetizado"}
	ENSINO_FUNDAMENTAL_INCOMPLETO EducationLevel = EducationLevel{"Ensino Fundamental incompleto"}
	ENSINO_FUNDAMENTAL_COMPLETO   EducationLevel = EducationLevel{"Ensino Fundamental completo"}
	ENSINO_MEDIO_INCOMPLETO       EducationLevel = EducationLevel{"Ensino Médio incompleto"}
	ENSINO_MEDIO_COMPLETO         EducationLevel = EducationLevel{"Ensino Médio completo"}
	ENSINO_SUPERIOR_INCOMPLETO    EducationLevel = EducationLevel{"Ensino Superior incompleto"}
	ENSINO_SUPERIOR_COMPLETO      EducationLevel = EducationLevel{"Ensino Superior completo"}
	POS_GRADUACAO                 EducationLevel = EducationLevel{"Pós-graduação"}
)
