package domains

type HomeType struct {
	homeType  string
}

const (
	ALUGADA = HomeType("ALUGADA")
	PARENTES = HomeType("PARENTES")
	FUNCIONAL = HomeType("FUNCIONAL")
	PROPRIA_QUITADA = HomeType("PROPRIA QUITADA")
	PROPRIA_FINANCIADA = HomeType("PROPRIA FINANCIADA")
	OUTROS = HomeType("OUTROS")
)