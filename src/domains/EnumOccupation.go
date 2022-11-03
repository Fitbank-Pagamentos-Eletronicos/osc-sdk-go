package domains

type Occupation struct {
	occupation string
}


const (
	APOSENTADO = Occupation{"Aposentado"}
	PENSIONISTA = Occupation{"Pensionista"}
	AUTONOMO = Occupation{"Autônomo"}
	EMPRESARIO = Occupation{"Empresário"}
	PROFISSIONAL_LIBERAL = Occupation{"Profissional Liberal"}
	ASSALARIOADO = Occupation{"Assalariado"}
	FUNCIONARIO_PUBLICO = Occupation{"Funcionário Público"}
	DESEMPREGADO = Occupation{"Desempregado"}
)