package entity

import (
	"gorm.io/gorm"
)

type ListenerModel struct {
	gorm.Model
	ip       string
	port     int32
	nickname string
}
