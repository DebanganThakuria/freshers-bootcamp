package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone int `json:"phone"`
	Address string `json:"address"`
}

func (u *User) TableName() string {
	return "user"
}