package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/entity"
	"github.com/indyspyz/web-service-market/service"
)

type MarketController interface {
	FindMarket() []entity.Market
	AddMarket(ctx *gin.Context) entity.Market
}

type marketController struct {
	service service.MarketService
}

func NewMarket(service service.MarketService) MarketController {
	return &marketController{
		service: service,
	}
}

func (c *marketController) FindMarket() []entity.Market {
	return c.service.FindMarket()
}

func (c *marketController) AddMarket(ctx *gin.Context) entity.Market {
	var market entity.Market
	ctx.BindJSON(&market)
	c.service.AddMarket(market)
	return market
}
