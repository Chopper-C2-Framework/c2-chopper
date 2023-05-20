package entity

type TaskModel struct {
	UUIDModel
	Name      string
	Repeat    bool
	AgentId   int
	Agent     AgentModel
	CreatorId int
	Creator   UserModel `gorm:"foreignKey:CreatorId"`
}
