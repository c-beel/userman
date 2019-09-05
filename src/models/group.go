package models

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name string `gorm:"unique"`
}
