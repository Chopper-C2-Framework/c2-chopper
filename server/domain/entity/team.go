package entity

import (
	"gorm.io/gorm"
)

type TeamModel struct {
	gorm.Model
	name    string
	members []*UserModel `gorm:"many2many:user_team;"`
}
