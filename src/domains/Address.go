package domains

type Address struct {
	zipCode string
	address string
	number string
	complement string
	district string
	state HomeTownState
	city string
	homeType HomeType
	homeSince HomeSince
}
