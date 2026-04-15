package main

import (
	"context"

	"github.com/Racemir/product-app/common/app"
	"github.com/Racemir/product-app/common/postgresql"
	"github.com/Racemir/product-app/controller"
	"github.com/Racemir/product-app/persistence"
	"github.com/Racemir/product-app/service"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)
	
	//web server başlatıyorum
	e.Start("localhost:8080")
}
