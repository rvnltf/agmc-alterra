package models

import (
	"github.com/jinzhu/gorm"
)

type Id struct {
	Id int64 `param:"id" json:"id"`
}
type Users struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	gorm.Model
}
type User struct {
	Fullname  string `json:"fullname" form:"fullname"`
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	CreatedAt string `json:"createdAt" form:"createdAt"`
	UpdatedAt string `json:"updatedAt" form:"updatedAt"`
}
