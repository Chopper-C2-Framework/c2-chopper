package entity

import (
	"time"
)

type AgentModel struct {
	UUIDModel
	Nickname string
	LastSeen time.Time
	Username string
	Uid      string
	Hostname string
}
