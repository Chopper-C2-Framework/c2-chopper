package entity

import "github.com/google/uuid"

type PluginResultModel struct {
	UUIDModel
	Path       string
	Output     string
	OutputType string
	CreatorId  uuid.UUID `type:"uuid"`
	Creator    UserModel `gorm:"foreignKey:CreatorId"`
}
