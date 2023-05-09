package entity

import (
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	Name     string
	Repeat   bool
	Listener ListenerModel `gorm:"embedded"`
}
