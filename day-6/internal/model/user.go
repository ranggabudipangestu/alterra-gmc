package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
