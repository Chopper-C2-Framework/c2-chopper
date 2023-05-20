package services

import (
	"fmt"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type TeamService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewTeamService(db *orm.ORMConnection) TeamService {
	logger := log.New()

	repo := entity.NewGormRepository(db.Db, logger)

	return TeamService{
		repo: repo,
	}
}

// CreateTeam Saving the new team to database
func (t TeamService) CreateTeam(newTeam *entity.TeamModel) error {

	fmt.Println("[+] Creating team", newTeam)
	err := t.repo.Create(newTeam)

	if err != nil {
		log.Debugf("failed to create team: %v\n", err)
		return err
	}

	return nil
}

func (t TeamService) FindOne(id string) (*entity.TeamModel, error) {
	var team entity.TeamModel
	err := t.repo.GetOneByID(&team, id)
	if err != nil {
		log.Debugf("Team not found\n")
		return nil, err

	}
	return &team, err
}

func (t TeamService) AddMemberToTeam(teamId string, userId string) error {
	var targetTeam *entity.TeamModel
	err := t.repo.GetOneByID(targetTeam, teamId)

	if err != nil {
		log.Debugf("failed to get team: %v\n", err)
		return err
	}

	var currentUser *entity.UserModel

	err = t.repo.GetOneByID(&currentUser, userId)

	if err != nil {
		log.Debugf("failed to get user: %v\n", err)
		return err
	}

	targetTeam.Members = append(targetTeam.Members, currentUser)

	err = t.repo.Save(targetTeam)

	if err != nil {
		log.Debugf("Error updating team %v", err)
		return err
	}

	return nil
}

func (t TeamService) UpdateTeam(toUpdateTeamId string, toUpdateTeam *entity.TeamModel) (*entity.TeamModel, error) {
	var targetTeam entity.TeamModel
	err := t.repo.GetOneByID(&targetTeam, toUpdateTeamId)

	if err != nil {
		log.Debugf("failed to get team for update: %v", err)
		return nil, err
	}

	targetTeam.Name = toUpdateTeam.Name
	targetTeam.Members = toUpdateTeam.Members

	err = t.repo.Save(&targetTeam)
	if err != nil {
		log.Debugf("failed to update team: %v", err)
		return nil, err
	}

	return &targetTeam, nil

}

func (t TeamService) DeleteTeam(teamId string) error {

	err := t.repo.Delete(&entity.TeamModel{
		UUIDModel: entity.UUIDModel{ID: uuid.MustParse(teamId)},
	})

	if err != nil {
		return err
	}

	return nil
}

func (t TeamService) GetAll() ([]entity.TeamModel, error) {
	var teams []entity.TeamModel
	err := t.repo.GetAll(&teams)
	if err != nil {
		return teams, err
	}

	return teams, nil
}
