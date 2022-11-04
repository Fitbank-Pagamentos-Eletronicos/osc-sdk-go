package domains

type AccountType struct {
	accountType  string
}

const (
	CONTA_CORRENTE_INDIVIDUAL AccountType = AccountType {"CONTA CORRENTE INDIVIDUAL"}
	CONTA_CORRENTE_CONJUNTA AccountType = AccountType{"CONTA CORRENTE CONJUNTA"}
	CONTA_POUPANCA_CONJUNTA AccountType  = AccountType{"CONTA POUPANCA CONJUNTA"}
	CONTA_POUPANCA_INDIVIDUAL  AccountType = AccountType{"CONTA POUPANCA INDIVIDUAL"}
)