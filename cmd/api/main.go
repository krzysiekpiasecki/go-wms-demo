package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kpiasecki/wms/internal/config"
	"github.com/kpiasecki/wms/internal/handler"
	"github.com/kpiasecki/wms/internal/logger"
	"github.com/kpiasecki/wms/internal/middleware"
	"github.com/kpiasecki/wms/internal/repository/postgres"
	"github.com/kpiasecki/wms/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewDatabase(
		cfg,
	)

	if err != nil {
		logger.Log.Fatal().
			Err(err).
			Msg("failed to connect to database")
	}
	defer db.Close(context.Background())

	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("/health", handler.Health)

	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	// Product handler
	productRepository := postgres.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	router.GET("/products/:id", productHandler.GetProduct)

	// Order handler
	orderRepository := postgres.NewOrderRepository(db)
	inventoryRepository := postgres.NewInventoryRepository(db)
	orderItemRepository := postgres.NewOrderItemRepository(db)

	orderService := service.NewOrderService(
		orderRepository,
		orderItemRepository,
		inventoryRepository,
	)

	orderHandler := handler.NewOrderHandler(orderService)
	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrder)
	router.PATCH("/orders/:id/status", orderHandler.UpdateStatus)

	err = router.Run(":8080")

	if err != nil {
		logger.Log.Fatal().
			Err(err).
			Msg("failed to start server")
	}

	logger.Log.Info().
		Str("component", "api").
		Msg("application started")

}
