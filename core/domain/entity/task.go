package entity

import "database/sql/driver"

type TaskType string

const (
	Ping      TaskType = "ping"
	Shell     TaskType = "shell"
	LiveShell TaskType = "live-shell"
)

type TaskModel struct {
	UUIDModel
	Name      string
	Type      TaskType `json:"type" sql:"type:ENUM('ping', 'shell', 'live-shell')"`
	AgentId   int
	Agent     AgentModel
	CreatorId int
	Creator   UserModel `gorm:"foreignKey:CreatorId"`
}

func (e *TaskType) Scan(value interface{}) error {
	*e = TaskType(value.([]byte))
	return nil
}

func (e TaskType) Value() (driver.Value, error) {
	return string(e), nil
}
