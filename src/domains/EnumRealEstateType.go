package domains

type RealEstateType struct {
	realEstateType string
}

var (
	HOUSE      RealEstateType = RealEstateType{"HOUSE"}
	APARTMENT  RealEstateType = RealEstateType{"APARTMENT"}
	COMMERCIAL RealEstateType = RealEstateType{"COMMERCIAL"}
	LAND       RealEstateType = RealEstateType{"LAND"}
	OTHERS     RealEstateType = RealEstateType{"OTHERS"}
)
