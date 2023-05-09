package entity

import (
	"gorm.io/gorm"
)

type TeamModel struct {
	gorm.Model
	name string
}
