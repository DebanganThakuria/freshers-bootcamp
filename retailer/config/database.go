package config

import (
	"fmt"
	"freshers-bootcamp/retailer/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildDBConfig(name string) *DBConfig {
	dbconfig := DBConfig{
		"localhost",
		3306,
		"root",
		name,
		"",
	}

	return &dbconfig
}

func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func GetDatabase(name string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(DBURL(BuildDBConfig(name))), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}