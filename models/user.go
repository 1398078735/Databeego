package models

type User struct {
	Phone string `form:"phone"`
	Pwd string `form:"password"`
}
