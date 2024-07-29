package estimate

import "cabbie/models"

type Response struct {
	Prices []Price `json:"prices"`
}

type Price struct {
	ServiceType models.ServiceTypeEnum `json:"serviceType"`
	Fare        float64                `json:"fare"`
}
