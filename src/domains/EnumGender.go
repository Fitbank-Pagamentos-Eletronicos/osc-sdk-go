package domains

type Gender struct {
	gender string
}

var (
	MASCULINO Gender = Gender{"MASCULINO"}
	FEMININO  Gender = Gender{"FEMININO"}
)
