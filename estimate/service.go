package estimate

import "math"

type Service struct {
	FareMultiplier map[ServiceTypeEnum]float64
}

func NewService() Service {
	fareMultiplier := map[ServiceTypeEnum]float64{
		BIKE:  20.0,
		CAR:   30.0,
		SEDAN: 40.0,
		SUV:   50.0,
	}
	return Service{FareMultiplier: fareMultiplier}
}

func (svc Service) ServeEstimate(estimateReq Request) (Response, error) {
	var prices []Price
	distance := math.Abs(estimateReq.DropOff.XCoordinate - estimateReq.Pickup.XCoordinate)
	for serviceType := BIKE; serviceType <= SUV; serviceType++ {
		prices = append(prices, Price{
			ServiceType: serviceType,
			Fare:        svc.FareMultiplier[serviceType] * distance,
		})
	}
	return Response{
		Prices: prices,
	}, nil
}
