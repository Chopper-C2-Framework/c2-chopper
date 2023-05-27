package services

import (
	"errors"

	log "github.com/sirupsen/logrus"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

	"time"
)

type AgentService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewAgentService(db *orm.ORMConnection) *AgentService {
	logger := log.New()

	repo := entity.NewGormRepository(db.Db, logger)

	return &AgentService{
		repo: repo,
	}
}

func (s *AgentService) FindAgentOrError(id string) (*entity.AgentModel, error) {
	var agent entity.AgentModel
	err := s.repo.GetOneByID(&agent, id)

	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (s *AgentService) CreateAgent(agent *entity.AgentModel) error {
	return s.repo.Create(agent)
}

func (s *AgentService) ConnectAgent(id string, agentInfo *entity.AgentModel) (*entity.AgentModel, error) {
	var (
		agent *entity.AgentModel
		err   error
	)

	if len(id) != 0 {
		agent, err = s.FindAgentOrError(id)
	} else {
		err = errors.New("Create new agent")
	}

	if err != nil {
		agent = &entity.AgentModel{}
		err = s.CreateAgent(agent)
	}
	if err != nil {
		return nil, err
	}
	agentInfo.LastSeen = time.Now()
	err = s.UpdateAgent(agent, agentInfo)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (s *AgentService) SaveAgent(agent *entity.AgentModel) error {
	return s.repo.Save(agent)
}

func (s *AgentService) UpdateAgent(target *entity.AgentModel, updates *entity.AgentModel) error {
	return s.repo.Update(target, updates)
}

func (s *AgentService) FindAllAgents() ([]*entity.AgentModel, error) {
	var res []*entity.AgentModel

	err := s.repo.GetAll(&res)
	if err != nil {
		log.Debugf("[-] failed to fetch agents")
		return nil, err
	}
	return res, nil
}
