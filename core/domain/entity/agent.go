package entity

type AgentModel struct {
	UUIDModel
	Ip       string
	Port     int32
	Nickname string
}
