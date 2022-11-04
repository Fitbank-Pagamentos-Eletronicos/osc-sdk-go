package domains

type NetWork struct {
	network string
}

var (
	VISA       NetWork = NetWork{"VISA"}
	MASTERCARD NetWork = NetWork{"MASTERCARD"}
	ELO        NetWork = NetWork{"ELO"}
)
