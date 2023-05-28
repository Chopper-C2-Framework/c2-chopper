package entity

type UserModel struct {
	UUIDModel
	Username string
	Password string
	Teams    []TeamModel `gorm:"many2many:team_user;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:team_id;jointable_foreignkey:user_model_id;"`
	Role     string      `gorm:"default:User"`
}
