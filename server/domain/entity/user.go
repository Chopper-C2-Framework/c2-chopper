package entity

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	username string
	password string
	teams    []*TeamModel `gorm:"many2many:user_team;"`
}
