package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type TeamManager struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewTeamService(db *orm.ORMConnection) *TeamManager {
	logger := log.New()
	repo := entity.NewGormRepository(db.Db, logger, "")
	instance := &entity.TeamModel{}

	if err := repo.Create(instance); err != nil {
		logger.Fatalf("failed to create cache instance: %v", err)
	}
	return &TeamManager{
		repo: repo,
	}
}

// Impelementation of the TeamManager interface with the orm package
func (t *TeamManager) CreateTeam(newTeam *entity.TeamModel) error {

	err := t.repo.Create(newTeam)

	if err != nil {
		log.Debugf("failed to create team: %v", err)
		return err
	}

	return nil
}

func (t *TeamManager) AddMemberToTeam(teamId string, user_id string) error {
	var targetTeam *entity.TeamModel
	err := t.repo.GetOneByID(targetTeam, teamId)

	if err != nil {
		log.Debugf("failed to get team: %v", err)
		return err
	}

	var currentUser *entity.UserModel

	err = t.repo.GetOneByID(&currentUser, user_id)

	if err != nil {
		log.Debugf("failed to get user: %v", err)
		return err
	}

	targetTeam.Members = append(targetTeam.Members, currentUser)

	err = t.repo.Save(targetTeam)

	return nil
}

func (t *TeamManager) UpdateTeam(toUpdateTeamId string, toUpdateTeam *entity.TeamModel) error {
	var targetTeam *entity.TeamModel
	err := t.repo.GetOneByID(targetTeam, toUpdateTeamId)

	if err != nil {
		log.Debugf("failed to get team for update: %v", err)
		return err
	}

	targetTeam.Name = toUpdateTeam.Name
	targetTeam.Members = toUpdateTeam.Members

	err = t.repo.Save(targetTeam)
	if err != nil {
		log.Debugf("failed to update team: %v", err)
		return err
	}

	return nil

}

func (t *TeamManager) DeleteTeam(team_id string) error {

	err := t.repo.Delete(&entity.TeamModel{
		UUIDModel: entity.UUIDModel{ID: uuid.MustParse(team_id)},
	})

	if err != nil {
		return err

	}

	return nil
}