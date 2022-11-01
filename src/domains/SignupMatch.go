package domains


type SignupMatch struct {
	cpf            string
	name           string
	birthday       string
	phone          string
	zipCode        string
	educationLevel []string
	banks          []Banks
	occupation     []string
	income         int
	hasCreditCart  bool
	hasRestriction bool
	hasOwnHouse    bool
	hasVehicle     bool
	hasAndroid     bool
}

func OSignupMatch(cpf string) string {

	var signupMatch SignupMatch

	signupMatch.educationLevel = []string{"NAO_ALFABETIZADO", "ENSINO_FUNDAMENTAL_INCOMPLETO", "ENSINO_FUNDAMENTAL_COMPLETO", "ENSINO_MEDIO_INCOMPLETO", "ENSINO_MEDIO_COMPLETO", "ENSINO_SUPERIOR_INCOMPLETO", "ENSINO_SUPERIOR_COMPLETO", "POS_GRADUACAO"}
	


	signupMatch.banks = []string{}

}
