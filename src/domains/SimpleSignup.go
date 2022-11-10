package domains

type SimpleSignup struct {
	Cpf            string
	Name           string
	Birthday       string
	Email          string
	Phone          string
	ZipCode        string
	HasCreditCart  bool
	HasRestriction bool
	HasOwnHouse    bool
	HasVehicle     bool
	HasAndroid     bool
	LogaData       LogData
}
