package models

import (
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) ([]Product, error) {
	products := []Product{}

	err := db.Order("id").Find(&products).Error

	return products, err
}

func GetProductByID(db *gorm.DB, id string) (Product, error) {
	product := Product{}

	err := db.Where("id = ?", id).First(&product).Error

	return product, err
}

func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

func UpdateProduct(db *gorm.DB, id string, price uint, quantity uint) (Product, error) {
	product := Product{}

	err := db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return product, err
	}

	product.Price = price
	product.Quantity = quantity

	db.Save(&product)

	return product, err
}