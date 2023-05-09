package entity

import (
	"gorm.io/gorm"

	"time"
)

type TaskResultModel struct {
	gorm.Model
	executedAt time.Time
	status     int32
}
