package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Nickname  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	FirstName string
	LastName  string
}
