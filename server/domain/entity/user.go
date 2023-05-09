package entity

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	username string
	password string
}
