package services

import (
	"fmt"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
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

func (t PluginResultService) GetPluginResultsOrError(path string) ([]*entity.PluginResultModel, error) {
	var results []*entity.PluginResultModel

	err := t.repo.GetByField(&results, "path", path)
	if err != nil {
		return nil, err
	}
	return results, nil
}
