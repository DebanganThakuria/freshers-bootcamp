package models

import (
	"gorm.io/gorm"
	"time"
)

func GetOrders(db *gorm.DB) ([]Order, error) {
	orders := []Order{}

	err := db.Order("id").Find(&orders).Error

	return orders, err
}

func CreateOrder(db *gorm.DB, order *Order, dict *map[uint]time.Time) string {
	cusID := order.CustomerID
	proID := order.ProductID
	OrderMap := *dict

	if 2 * time.Minute > time.Now().Sub(OrderMap[cusID]) {
		return "Failed. Wait for some time"
	}

	product := Product{}
	err := db.Where("id = ?", proID).Find(&product).Error
	if err != nil {
		return "failed"
	}

	if product.Quantity < order.Quantity || product.Quantity == 0 {
		return "failed"
	}

	err = db.Create(order).Error
	if err != nil {
		return "failed"
	}

	product.Quantity -= order.Quantity
	db.Save(&product)
	OrderMap[cusID] = time.Now()

	return "order placed"
}

func GetOrderByID(db *gorm.DB, id string) (Order, bool) {
	order := Order{}

	err := db.Where("id = ?", id).Find(&order).Error
	if err != nil {
		return order, false
	}
	return order, true
}