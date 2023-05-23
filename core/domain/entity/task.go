package entity

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type TaskType string

const (
	Ping  TaskType = "ping"
	Shell TaskType = "shell"
)

type TaskModel struct {
	UUIDModel
	Name      string
	Type      TaskType `json:"type" sql:"type:ENUM('ping', 'shell')"`
	Args      string
	AgentId   uuid.UUID
	Agent     AgentModel
	CreatorId uuid.UUID
	Creator   UserModel `gorm:"foreignKey:CreatorId"`
}

func (e *TaskType) Scan(value interface{}) error {
	*e = TaskType(value.([]byte))
	return nil
}

func (e TaskType) Value() (driver.Value, error) {
	return string(e), nil
}
