package service

import (
	"github.com/indyspyz/web-service-market/entity"
)

type MarketService interface {
	AddMarket(entity.Market) entity.Market
	FindMarket() []entity.Market
}

type marketService struct {
	markets []entity.Market
}

func NewMarket() MarketService {
	return &marketService{}
}

func (service *marketService) AddMarket(market entity.Market) entity.Market {
	service.markets = append(service.markets, market)
	return market
}

func (service *marketService) FindMarket() []entity.Market {
	return service.markets
}
