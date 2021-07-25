package routes

import (
	"freshers-bootcamp/first-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	group1 := r.Group("/user-api")
	{
		group1.GET("user", controllers.GetUsers)
		group1.POST("user", controllers.CreateUser)
		group1.GET("user/:id", controllers.GetUserByID)
		group1.PUT("user/:id", controllers.UpdateUser)
		group1.DELETE("user/:id", controllers.DeleteUser)
	}

	return r
}