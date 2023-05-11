package entity

import (
	"gorm.io/gorm"

	"time"
)

type TaskResultModel struct {
	gorm.Model
	ExecutedAt time.Time
	Status     int32
	TaskID     int
	Task       *TaskModel
}
