package entity

import (
	"time"
)

type TaskResultModel struct {
	UUIDModel
	ExecutedAt time.Time
	Status     int32
	TaskID     int
	Task       *TaskModel
}
