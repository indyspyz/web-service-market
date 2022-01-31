package service

import (
	"github.com/indyspyz/web-service-market/entity"
)

type MarketService interface {
	Save(entity.Market) entity.Market
	FindAll() []entity.Market
}

type marketService struct {
	markets []entity.Market
}

func New() MarketService {
	return &marketService{}
}

func (service *marketService) Save(market entity.Market) entity.Market {
	service.markets = append(service.markets, market)
	return market
}

func (service *marketService) FindAll() []entity.Market {
	return service.markets
}
