package models

type Customer struct {
	ID uint `json:"id"`
	Password string `json:"password"`
}

func (u *Customer) TableName() string {
	return "customer"
}