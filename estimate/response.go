package estimate

type Response struct {
	Prices []Price `json:"prices"`
}

type Price struct {
	ServiceType ServiceTypeEnum `json:"serviceType"`
	Fare        float64         `json:"fare"`
}

type ServiceTypeEnum int

const (
	UNKNOWN ServiceTypeEnum = iota
	BIKE
	CAR
	SEDAN
	SUV
)
