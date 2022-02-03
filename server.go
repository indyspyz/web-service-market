package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/indyspyz/web-service-market/controller"
	"github.com/indyspyz/web-service-market/middlewares"
	"github.com/indyspyz/web-service-market/service"
)

var (
	marketService    service.MarketService       = service.NewMarket()
	MarketController controller.MarketController = controller.NewMarket(marketService)

	productService    service.ProductService       = service.NewProduct()
	ProductController controller.ProductController = controller.NewProduct(productService)
)

func setuoLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setuoLogOutPut()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/markets/findmarket", func(ctx *gin.Context) {
		ctx.JSON(200, MarketController.FindMarket())
	})

	server.POST("/markets/addmarket", func(ctx *gin.Context) {
		ctx.JSON(200, MarketController.AddMarket(ctx))
	})

	server.GET("/products/findproduct", func(ctx *gin.Context) {
		ctx.JSON(200, ProductController.FindProduct())
	})

	server.POST("/products/addproduct", func(ctx *gin.Context) {
		ctx.JSON(200, ProductController.AddProduct(ctx))
	})

	server.Run(":8080")
}
