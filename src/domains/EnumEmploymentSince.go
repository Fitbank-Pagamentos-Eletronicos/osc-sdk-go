package domains

type EmploymentSince struct {
	employmentSince  string
}

const (
	LESS_THAN_SIX_MONTHS = EmploymentSince("LESS THAN SIX MONTHS")
	BETWEEN_SIX_MONTHS_AND_ONE_YEAR = EmploymentSince("BETWEEN SIX MONTHS AND ONE YEAR")
	BETWEEN_ONE_AND_TWO_YEARS  = EmploymentSince("BETWEEN ONE AND TWO YEARS")
	BETWEEN_TWO_AND_THREE_YEARS  = EmploymentSince("BETWEEN TWO AND THREE YEARS")
	BETWEEN_THREE_AND_FOUR_YEARS  = EmploymentSince("BETWEEN THREE AND FOUR YEARS")
	BETWEEN_FOUR_AND_FIVE_YEARS  = EmploymentSince("BETWEEN FOUR AND FIVE YEARS")
	BETWEEN_FIVE_AND_TEN_YEARS = EmploymentSince("BETWEEN FIVE AND TEN YEARS")
	MORE_THAN_TEN_YEARS = EmploymentSince("MORE THAN TEN YEARS")
)
