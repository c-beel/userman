package models

import (
	"github.com/c-beel/userman/src/pkg/api/v1"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/cloudsearch/v1"
)

type User struct {
	gorm.Model
	username string
	nickname string
	email string
	firstName string
	lastName string
}
