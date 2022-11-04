package domains

type EmploymentSince struct {
	employmentSince string
}

var (
	LESS_THAN_SIX_MONTHS            EmploymentSince = EmploymentSince{"LESS THAN SIX MONTHS"}
	BETWEEN_SIX_MONTHS_AND_ONE_YEAR EmploymentSince = EmploymentSince{"BETWEEN SIX MONTHS AND ONE YEAR"}
	BETWEEN_ONE_AND_TWO_YEARS       EmploymentSince = EmploymentSince{"BETWEEN ONE AND TWO YEARS"}
	BETWEEN_TWO_AND_THREE_YEARS     EmploymentSince = EmploymentSince{"BETWEEN TWO AND THREE YEARS"}
	BETWEEN_THREE_AND_FOUR_YEARS    EmploymentSince = EmploymentSince{"BETWEEN THREE AND FOUR YEARS"}
	BETWEEN_FOUR_AND_FIVE_YEARS     EmploymentSince = EmploymentSince{"BETWEEN FOUR AND FIVE YEARS"}
	BETWEEN_FIVE_AND_TEN_YEARS      EmploymentSince = EmploymentSince{"BETWEEN FIVE AND TEN YEARS"}
	MORE_THAN_TEN_YEARS             EmploymentSince = EmploymentSince{"MORE THAN TEN YEARS"}
)
