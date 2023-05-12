package entity

type UserModel struct {
	UUIDModel
	Username string
	Password string
	Teams    []*TeamModel `gorm:"many2many:user_team_membership;"`
}
