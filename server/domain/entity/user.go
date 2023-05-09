package entity

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string
	Password string
	Teams    []*TeamModel `gorm:"many2many:user_team_membership;"`
}
