package models

import "github.com/jinzhu/gorm"

type Membership struct {
	gorm.Model
	User  User  `gorm:"foreignkey:UID"`
	UID   int64 `gorm:"column:uid"`
	Group Group `gorm:"foreignkey:GID"`
	GID   int64 `gorm:"column:gid"`
}
