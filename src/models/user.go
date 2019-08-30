package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	username  string
	nickname  string
	email     string
	firstName string
	lastName  string
}
