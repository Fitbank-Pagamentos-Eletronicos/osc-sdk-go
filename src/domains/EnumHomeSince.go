package domains

type HomeSince struct {
	homeSince string
}

var (
	MENOR_6_MESES HomeSince = HomeSince{"MENOR 6 MESES"}
	MENOR_1_ANO   HomeSince = HomeSince{"MENOR 1 ANO"}
	MENOR_2_ANOS  HomeSince = HomeSince{"MENOR 2 ANOS"}
	MAIOR_2_ANOS  HomeSince = HomeSince{"MAIOR 2 ANOS"}
)
