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