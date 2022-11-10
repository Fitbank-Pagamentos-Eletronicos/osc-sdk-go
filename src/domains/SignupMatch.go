package domains

type SignupMatch struct {
	Cpf            string
	Name           string
	Birthday       string
	Phone          string
	ZipCode        string
	EducationLevel EducationLevel
	Banks          Banks
	Occupation     Occupation
	Income         int
	HasCreditCart  bool
	HasRestriction bool
	HasOwnHouse    bool
	HasVehicle     bool
	HasAndroid     bool
	Products       Products
	LogData        LogData
}
