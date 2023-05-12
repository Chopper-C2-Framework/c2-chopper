package entity

type TaskModel struct {
	UUIDModel
	Name       string
	Repeat     bool
	ListenerId int
	Listener   ListenerModel
	CreatorId  int
	Creator    UserModel `gorm:"foreignKey:CreatorId"`
}
