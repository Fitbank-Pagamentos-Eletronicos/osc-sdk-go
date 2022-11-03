package domains

type SignupMatch struct {
	cpf            string
	name           string
	birthday       string
	phone          string
	zipCode        string
	educationLevel EducationLevel
	banks          Banks
	occupation     Occupation
	income         int
	hasCreditCart  bool
	hasRestriction bool
	hasOwnHouse    bool
	hasVehicle     bool
	hasAndroid     bool
}





