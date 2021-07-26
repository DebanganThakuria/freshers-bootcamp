package main

import (
	"fmt"
	"freshers-bootcamp/student-api/config"
	"freshers-bootcamp/student-api/models"
	"freshers-bootcamp/student-api/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DBURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Student{}, &models.Subject{}, &models.StudentMarks{})

	r := routes.SetUpRouter()
	r.Run()
}
