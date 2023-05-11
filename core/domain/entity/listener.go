package entity

import (
	"gorm.io/gorm"
)

type ListenerModel struct {
	gorm.Model
	Ip       string
	Port     int32
	Nickname string
}
