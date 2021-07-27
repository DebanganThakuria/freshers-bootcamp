package routes

import (
	"freshers-bootcamp/retailer/config"
	"freshers-bootcamp/retailer/handlers"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(name string) *gin.Engine {
	r := gin.Default()

	api := handlers.API{
		config.GetDatabase(name),
		make(map[uint]time.Time),
	}

	group1 := r.Group("/product")
	{
		group1.POST("/", api.CreateProduct)
		group1.PATCH("/:id", api.UpdateProduct)
		group1.GET("/:id", api.GetProductByID)
		group1.GET("/", api.GetProducts)
	}

	group2 := r.Group("/order")
	{
		group2.POST("/", api.CreateOrder) // multiple products in 1 order
		group2.GET("/", api.GetOrders)
		group2.GET("/:id", api.GetOrderByID)
	}

	group3 := r.Group("/customer")
	{
		group3.POST("/", api.CreateCustomer)
		group3.GET("/", api.GetCustomers)
	}

	return r
}