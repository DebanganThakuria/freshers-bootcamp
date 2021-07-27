package models

type Product struct {
	ID uint `json:"id"`
	ProductName string `json:"product_name"`
	Price uint `json:"price"`
	Quantity uint `json:"quantity"`
}

func (u *Product) TableName() string {
	return "product"
}