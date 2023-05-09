package entity

import (
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	Name       string
	Repeat     bool
	ListenerId int
	Listener   ListenerModel
}
