package domains

type Nationality struct {
	nationality string
}

var (
	BRASILEIRA  Nationality = Nationality{"BRASILEIRA"}
	ESTRANGEIRA Nationality = Nationality{"ESTRANGEIRA"}
)
