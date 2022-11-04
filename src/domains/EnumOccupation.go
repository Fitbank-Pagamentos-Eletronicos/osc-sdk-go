package domains

type Occupation struct {
	occupation string
}

var (
	APOSENTADO           Occupation = Occupation{"Aposentado"}
	PENSIONISTA          Occupation = Occupation{"Pensionista"}
	AUTONOMO             Occupation = Occupation{"Autônomo"}
	EMPRESARIO           Occupation = Occupation{"Empresário"}
	PROFISSIONAL_LIBERAL Occupation = Occupation{"Profissional Liberal"}
	ASSALARIOADO         Occupation = Occupation{"Assalariado"}
	FUNCIONARIO_PUBLICO  Occupation = Occupation{"Funcionário Público"}
	DESEMPREGADO         Occupation = Occupation{"Desempregado"}
)
