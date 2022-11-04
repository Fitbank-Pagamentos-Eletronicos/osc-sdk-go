package domains

type HomeType struct {
	homeType string
}

var (
	ALUGADA            HomeType = HomeType{"ALUGADA"}
	PARENTES           HomeType = HomeType{"PARENTES"}
	FUNCIONAL          HomeType = HomeType{"FUNCIONAL"}
	PROPRIA_QUITADA    HomeType = HomeType{"PROPRIA QUITADA"}
	PROPRIA_FINANCIADA HomeType = HomeType{"PROPRIA FINANCIADA"}
	OUTROS             HomeType = HomeType{"OUTROS"}
)
