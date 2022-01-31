package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/entity"
	"github.com/indyspyz/web-service-market/service"
)

type MarketController interface {
	FindAll() []entity.Market
	Save(ctx *gin.Context) entity.Market
}

type controller struct {
	service service.MarketService
}

func New(service service.MarketService) MarketController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Market {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Market {
	var market entity.Market
	ctx.BindJSON(&market)
	c.service.Save(market)
	return market
}
