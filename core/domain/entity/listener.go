package entity

type ListenerModel struct {
	UUIDModel
	Ip       string
	Port     int32
	Nickname string
}
