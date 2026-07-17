package main

import (
	"github.com/gin-contrib/cors"
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
	defer db.Close()

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4200",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
		},
	}))

	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("/health", handler.Health)

	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	inventoryRepository := postgres.NewInventoryRepository(db)

	// Product handler
	productRepository := postgres.NewProductRepository(db)
	productService := service.NewProductService(productRepository, inventoryRepository)
	productHandler := handler.NewProductHandler(productService)
	router.GET("/products/:id", productHandler.GetProduct)
	router.GET("/products", productHandler.GetProducts)
	router.POST("/products", productHandler.CreateProduct)

	// Order handler
	orderRepository := postgres.NewOrderRepository(db)

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
	router.GET("/orders", orderHandler.GetOrders)

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
