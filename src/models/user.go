package models

import (
	"github.com/c-beel/userman/src/pkg/api/v1"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	v1.User
}
