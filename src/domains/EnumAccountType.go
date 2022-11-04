package domains

type AccountType struct {
	accountType string
}

var (
	CONTA_CORRENTE_INDIVIDUAL AccountType = AccountType{"CONTA_CORRENTE_INDIVIDUAL"}
	CONTA_CORRENTE_CONJUNTA   AccountType = AccountType{"CONTA_CORRENTE_CONJUNTA"}
	CONTA_POUPANCA_CONJUNTA   AccountType = AccountType{"CONTA_POUPANCA"}
	CONTA_POUPANCA_INDIVIDUAL AccountType = AccountType{"CONTA_POUPANCA_INDIVIDUAL"}
)
