package models

type Order struct {
	ID uint `json:"id"`
	Quantity uint `json:"quantity"`
	CustomerID uint `json:"customer_id"`
	Customer Customer `gorm:"foreignKey:customer_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID uint `json:"product_id"`
	Product Product `gorm:"foreignKey:product_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *Order) TableName() string {
	return "order"
}