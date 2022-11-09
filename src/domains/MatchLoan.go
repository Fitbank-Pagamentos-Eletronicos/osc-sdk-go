package domains

type MatchLoan struct {
	ProductId      int
	Name           string
	Logo           string
	MinValue       int
	MaxValue       int
	MinInstallment int
	MaxInstallment int
	MonthlyTax     float64
}
