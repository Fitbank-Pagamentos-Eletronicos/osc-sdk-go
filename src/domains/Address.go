package domains

type Address struct {
	ZipCode string
	Address string
	Number string
	Complement string
	District string
	State HomeTownState
	City string
	HomeType_ HomeType
	HomeSince_ HomeSince
}
