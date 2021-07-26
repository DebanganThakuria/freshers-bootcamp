package routes

import (
	"freshers-bootcamp/student-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	group1 := r.Group("/student")
	{
		group1.GET("/", controllers.GetStudents)
		group1.POST("/", controllers.CreateStudent)
	}

	group2 := r.Group("/subject")
	{
		group2.GET("/", controllers.GetSubjects)
		group2.POST("/", controllers.CreateSubject)
	}

	group3 := r.Group("/student-marks")
	{
		group3.POST("/", controllers.CreateStudentMarks)
		group3.GET("/:id", controllers.GetStudentMarks)
	}

	return r
}
