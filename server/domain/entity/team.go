package entity

import (
	"gorm.io/gorm"
)

type TeamModel struct {
	gorm.Model
	Name    string
	Members []*UserModel `gorm:"many2many:user_team;"`
}
