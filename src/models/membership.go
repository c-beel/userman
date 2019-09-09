package models

import "github.com/jinzhu/gorm"

type Membership struct {
	gorm.Model
	User  User  `gorm:"foreignkey:UID"`
	UID   uint `gorm:"column:uid"`
	Group Group `gorm:"foreignkey:GID"`
	GID   uint `gorm:"column:gid"`
}
