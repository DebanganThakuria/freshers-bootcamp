package handlers

import (
	"fmt"
	"freshers-bootcamp/retailer/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *API) CreateCustomer(c *gin.Context) {
	customer := models.Customer{}

	c.BindJSON(&customer)

	err := models.CreateCustomer(a.DB, &customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func (a *API) GetCustomers(c *gin.Context) {
	customers, err := models.GetCustomers(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (a *API) CustomerHistory(c *gin.Context) {
	id := c.Params.ByName("id")
	orders, err := models.CustomerHistory(a.DB, id)
	if !err {
		c.JSON(http.StatusInternalServerError, "No History")
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