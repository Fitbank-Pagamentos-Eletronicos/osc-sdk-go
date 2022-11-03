package domains

type RealEstateType struct{
	realEstateType string
}

const (
	HOUSE = RealEstateType{"HOUSE"}
	APARTMENT = RealEstateType{"APARTMENT"}
	COMMERCIAL = RealEstateType{"COMMERCIAL"}
	LAND = RealEstateType{"LAND"}
	OTHERS = RealEstateType{"OTHERS"}
)