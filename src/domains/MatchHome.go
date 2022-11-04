package domains

type MatchHome struct {
	productId      int16
	name           string
	logo           string
	minValue       int8
	maxValue       int16
	minInstallment int16
	maxInstallment int16
	monthlyTax     float32
}
