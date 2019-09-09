package models

import (
	"github.com/jinzhu/gorm"
	"strings"
	"errors"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Nickname  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	FirstName string
	LastName  string
}

func (user User) Validate(db *gorm.DB) {
	if !strings.HasSuffix(user.Email, "@google.com") {
		db.AddError(errors.New("invalid email address domain"))
	}
	if user.Email == "" {
		db.AddError(errors.New("empty email address"))
	}
}
