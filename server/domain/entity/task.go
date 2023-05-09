package entity

import (
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model // Contains createdAt
	name       string
	repeat     bool
	listener   ListenerModel `gorm:"embedded"`
	// createdAt time.Time
}
