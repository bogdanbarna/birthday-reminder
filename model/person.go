package model

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Person struct {
	Username string `form:"username" gorm:"unique"`
	Birthday string `json:"birthday"`
	gorm.Model
}
