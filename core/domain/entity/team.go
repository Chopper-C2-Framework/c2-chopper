package entity

type TeamModel struct {
	UUIDModel
	Name    string
	Members []*UserModel `gorm:"many2many:user_team;"`
	// Hosts   []*HostModel `gorm:"many2many:host_team;"`
}

// All operations that can be performed on the TeamModel
// type TeamRepository struct {
// 	Db *gorm.DB
// }

// // Implement the interface, it uses gorm to create the
// // TeamModel
// func (t *TeamRepository) CreateTeam() error {
// 	return nil
// }

// func (t *TeamRepository) JoinTeam() error {
// 	return nil
// }

// func (t *TeamRepository) AddMemberToTeam() error {
// 	return nil
// }

// func (t *TeamRepository) UpdateTeam() error {
// 	return nil
// }

// func (t *TeamRepository) DeleteTeam() error {
// 	return nil
// }
