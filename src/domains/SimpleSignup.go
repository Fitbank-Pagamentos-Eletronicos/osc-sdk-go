package domains

type SimpleSignup struct {
	cpf            string
	name		   string
	birthday	   string
	email		   string
	phone		   string
	zipCode		   string
	hasCreditCart  bool
	hasRestriction bool
	hasOwnHouse    bool
	hasVehicle     bool
	hasAndroid     bool
	logaData 	 LogData

}