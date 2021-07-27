package handlers

import (
	"fmt"
	"freshers-bootcamp/retailer/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderWrapper struct {
	OrderID uint
	CustomerID uint
	ProductID uint
	Quantity uint
	Status string
}

func (a *API) GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	order, err := models.GetOrderByID(a.DB, id)
	if !false {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "order not found")
		return
	}

	c.JSON(http.StatusOK, OrderWrapper{
			order.ID,
			order.CustomerID,
			order.ProductID,
			order.Quantity,
			"processed",
		})
}

func (a *API) CreateOrder(c *gin.Context) {
	order := models.Order{}

	c.BindJSON(&order)

	status := models.CreateOrder(a.DB, &order, &a.OrderMap)

	if status == "order placed" {
		c.JSON(http.StatusOK, OrderWrapper{
			order.ID,
			order.CustomerID,
			order.ProductID,
			order.Quantity,
			status,
		})
		return
	}
	c.JSON(http.StatusNotFound, status)
}

func (a *API) GetOrders(c *gin.Context) {
	orders, err := models.GetOrders(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ordersResponse := []OrderWrapper{}

	for _, val := range orders {
		wrapper := OrderWrapper{
			val.ID,
			val.CustomerID,
			val.ProductID,
			val.Quantity,
			"processed",
		}
		ordersResponse = append(ordersResponse, wrapper)
	}

	c.JSON(http.StatusOK, ordersResponse)
}