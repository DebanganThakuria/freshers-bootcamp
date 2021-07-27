package models

import "gorm.io/gorm"

func GetCustomers(db *gorm.DB) ([]Customer, error) {
	customers := []Customer{}

	err := db.Order("id").Find(&customers).Error

	return customers, err
}

func CreateCustomer(db *gorm.DB, customer *Customer) error {
	return db.Create(customer).Error
}