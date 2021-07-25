package main

import (
	"fmt"
	"freshers-bootcamp/first-api/config"
	"freshers-bootcamp/first-api/models"
	"freshers-bootcamp/first-api/routes"
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

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetUpRouter()
	r.Run()
}