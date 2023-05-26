package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskResultModel struct {
	UUIDModel
	ExecutedAt time.Time
	Status     int32
	TaskID     uuid.UUID `type:"uuid"`
	Task       *TaskModel
	Output     string
	Seen       bool `gorm:"default:false"`
}
