package estimate

type Service struct{}

func NewService() Service {
	return Service{}
}

func (svc Service) ServeEstimate(estimateReq Request) (Response, error) {
	return Response{
		Prices: []Price{
			{
				ServiceType: 1,
				Fare:        float64(50),
			},
		},
	}, nil
}
