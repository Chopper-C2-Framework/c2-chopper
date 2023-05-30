package services

import (
	"fmt"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type PluginResultService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewPluginResultService(db *orm.ORMConnection) PluginResultService {
	logger := log.New()

	repo := entity.NewGormRepository(db.Db, logger)
	return PluginResultService{
		repo: repo,
	}
}

// CreateTeam Saving the new team to database
func (t PluginResultService) CreatePluginResult(newRes *entity.PluginResultModel) error {
	fmt.Println("[+] Creating plugin result", newRes)
	err := t.repo.Create(newRes)

	if err != nil {
		log.Debugf("failed to create plugin result: %v\n", err)
		return err
	}
	return nil
}

func (t PluginResultService) DeletePluginResult(id string) error {
	res, err := t.FindResultByIdOrError(id)
	if err != nil {
		return err
	}

	err = t.repo.Delete(res)
	if err != nil {
		return err
	}
	return nil
}

func (t PluginResultService) FindResultByIdOrError(id string) (*entity.PluginResultModel, error) {
	var res entity.PluginResultModel
	err := t.repo.GetOneByID(&res, id)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (t PluginResultService) GetPluginResultsOrError(path string, userId uuid.UUID) ([]*entity.PluginResultModel, error) {
	var results []*entity.PluginResultModel
	var filters map[string]interface{} = make(map[string]interface{})

	filters["path"] = path
	filters["creator_id"] = userId

	err := t.repo.GetByFields(&results, filters)
	if err != nil {
		return nil, err
	}
	return results, nil
}
