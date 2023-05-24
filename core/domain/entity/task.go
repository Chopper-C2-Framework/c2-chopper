package entity

import (
	// "database/sql/driver"

	"github.com/google/uuid"
)

type TaskType string

const (
	TASK_TYPE_PING  TaskType = "PING"
	TASK_TYPE_SHELL TaskType = "SHELL"
)

type TaskModel struct {
	UUIDModel
	Name      string
	Type      TaskType `json:"type" sql:"type:ENUM('PING', 'SHELL')"`
	Args      string
	AgentId   uuid.UUID `type:"uuid"`
	Agent     AgentModel
	CreatorId uuid.UUID `type:"uuid"`
	Creator   UserModel `gorm:"foreignKey:CreatorId"`
}

// func (e *TaskType) Scan(value interface{}) error {
// 	*e = TaskType(value.([]byte))
// 	return nil
// }

// func (e TaskType) Value() (driver.Value, error) {
// 	return string(e), nil
// }
