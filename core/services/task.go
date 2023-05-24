package services

import (
	log "github.com/sirupsen/logrus"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type TaskService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewTaskService(db *orm.ORMConnection) *TaskService {
	logger := log.New()

	repo := entity.NewGormRepository(db.Db, logger)
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(task *entity.TaskModel) error {
	err := s.repo.Create(task)
	if err != nil {
		log.Debugf("[-] failed to create task")
		return err
	}

	return nil
}

func (s *TaskService) DeleteTask(task *entity.TaskModel) error {
	err := s.repo.Delete(task)
	if err != nil {
		log.Debugf("[-] failed to delete task")
		return err
	}

	return nil
}

func (s *TaskService) FindTaskOrError(taskId string) (*entity.TaskModel, error) {
	var task entity.TaskModel
	err := s.repo.GetOneByID(&task, taskId)

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) FindTasksForAgent(agentId string) ([]entity.TaskModel, error) {
	var tasks []entity.TaskModel

	err := s.repo.GetByField(&tasks, "agent_id", agentId)
	if err != nil {
		log.Debugf("[-] failed to find task by agentid")
		return nil, err
	}
	return tasks, nil
}
