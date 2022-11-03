package domains

type NetWork struct{
	network string
}

const (
	VISA = NetWork{"VISA"}
	MASTERCARD = NetWork{"MASTERCARD"}
	ELO = NetWork{"ELO"}
)