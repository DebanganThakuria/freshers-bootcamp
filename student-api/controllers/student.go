package controllers

import (
	"fmt"
	"freshers-bootcamp/student-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Students Table C and R Operations
func GetStudents(c *gin.Context) {
	var student []models.Student

	err := models.GetAllStudents(&student)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	c.BindJSON(&student)

	err := models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// Subjects Table C and R Operations
func GetSubjects(c *gin.Context) {
	var subject []models.Subject

	err := models.GetAllSubjects(&subject)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}

func CreateSubject(c *gin.Context) {
	var subject models.Subject

	c.BindJSON(&subject)

	err := models.CreateSubject(&subject)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}

// StudentsMarks Table CRUD Operations
func CreateStudentMarks(c *gin.Context) {
	var marks models.StudentMarks

	c.BindJSON(&marks)
	err := models.CreateStudentMarks(&marks)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

func GetStudentMarks(c *gin.Context) {
	studentID := c.Params.ByName("id")

	var marks []models.StudentMarks

	err := models.GetStudentMarksByID(&marks, studentID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}
