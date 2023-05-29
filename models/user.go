package models

type User struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
