package handlers

import (
	"fmt"
	"freshers-bootcamp/retailer/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductWrapper struct {
	Product models.Product
	Message string
}

type ProductUpdate struct {
	Price uint `json:"price"`
	Quantity uint `json:"quantity"`
}

func (a *API) GetProducts(c *gin.Context) {
	products, err := models.GetProducts(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (a *API) CreateProduct(c *gin.Context) {
	product := models.Product{}

	c.BindJSON(&product)

	err := models.CreateProduct(a.DB, &product)
	if err != nil {
		c.JSON(http.StatusNotFound, ProductWrapper{product, "Failed"})
	} else {
		c.JSON(http.StatusOK, ProductWrapper{product, "product successfully added"})
	}
}

func (a *API) UpdateProduct(c *gin.Context) {
	update := ProductUpdate{}

	c.BindJSON(&update)
	id := c.Params.ByName("id")

	product, err := models.UpdateProduct(a.DB, id, update.Price, update.Quantity)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func (a *API) GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")

	product, err := models.GetProductByID(a.DB, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}