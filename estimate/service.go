package estimate

import (
	"cabbie/models"
	"math"
)

type Service struct {
	FareMultiplier map[models.ServiceTypeEnum]float64
}

func NewService() Service {
	fareMultiplier := map[models.ServiceTypeEnum]float64{
		models.BIKE:  20.0,
		models.CAR:   30.0,
		models.SEDAN: 40.0,
		models.SUV:   50.0,
	}
	return Service{FareMultiplier: fareMultiplier}
}

func (svc Service) ServeEstimate(estimateReq Request) (Response, error) {
	var prices []Price
	// simple estimate logic to calculate fare as a multiple of distance, as described by the fare multiplier stored with the service instance
	// this can be easily extended to be configurable, or an API can be exposed to modify this
	distance := math.Abs(estimateReq.Destination.XCoordinate - estimateReq.Origin.XCoordinate)
	for serviceType := models.BIKE; serviceType <= models.SUV; serviceType++ {
		prices = append(prices, Price{
			ServiceType: serviceType,
			Fare:        svc.FareMultiplier[serviceType] * distance,
		})
	}
	return Response{
		Prices: prices,
	}, nil
}
