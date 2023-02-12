package service

type Service struct {
	Rates *RatesService
}

func NewService() *Service {
	return &Service{
		Rates: NewRatesService(),
	}
}
