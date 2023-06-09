package services

import (
	log "github.com/sirupsen/logrus"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
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

func (s *TaskService) UpdateTask(target *entity.TaskModel, updates *entity.TaskModel) error {
	return s.repo.Update(target, updates)
}

func (s *TaskService) FindTaskOrError(taskId string) (*entity.TaskModel, error) {
	var task entity.TaskModel
	err := s.repo.GetOneByID(&task, taskId)

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) FindTasksForAgent(agentId string) ([]*entity.TaskModel, error) {
	var tasks []*entity.TaskModel

	err := s.repo.GetByField(&tasks, "agent_id", agentId)
	if err != nil {
		log.Debugf("[-] failed to find task by agentid")
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) FindAllTasks() ([]*entity.TaskModel, error) {
	var tasks []*entity.TaskModel

	err := s.repo.GetAll(&tasks)
	if err != nil {
		log.Debugf("[-] failed to fetch tasks")
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) FindUnexecutedTasks() ([]*entity.TaskModel, error) {
	var tasks []*entity.TaskModel

	err := s.repo.DB().
		Where("NOT EXISTS(SELECT id FROM task_result_models WHERE task_id = task_models.id)").
		Find(&tasks).Error
	if err != nil {
		log.Debugf("[-] failed to fetch tasks")
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) FindNewlyExecutedTasks() ([]*entity.TaskModel, error) {
	var tasks []*entity.TaskModel

	err := s.repo.DB().
		Where("EXISTS(SELECT id FROM task_result_models WHERE task_id = task_models.id AND seen=false)").
		Find(&tasks).Error
	if err != nil {
		log.Debugf("[-] failed to fetch tasks")
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) FindUnexecutedTasksForAgent(agentId string) ([]*entity.TaskModel, error) {
	var tasks []*entity.TaskModel

	err := s.repo.DB().
		Where("agent_id = ?", agentId).
		Where("NOT EXISTS(SELECT id FROM task_result_models WHERE task_id = task_models.id)").
		Find(&tasks).Error

	if err != nil {
		log.Debugf("[-] failed to find task by agentid")
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) CreateTaskResult(taskResult *entity.TaskResultModel) error {
	err := s.repo.Create(taskResult)
	if err != nil {
		log.Debugf("[-] failed to create task")
		return err
	}

	return nil
}

func (s *TaskService) FindTaskResults(taskId string) ([]*entity.TaskResultModel, error) {
	var res []*entity.TaskResultModel

	err := s.repo.GetByField(&res, "task_id", taskId)
	if err != nil {
		log.Debugf("[-] failed to find task results by taskid")
		return nil, err
	}
	return res, nil
}

func (s *TaskService) FindLatestResults(limit uint32, page uint32) ([]*entity.TaskResultModel, error) {
	var res []*entity.TaskResultModel

	err := s.repo.GetSortedBatch(&res, "created_at", true, int(limit), int(page*limit))
	if err != nil {
		log.Debugf("[-] failed to find task results by taskid")
		return nil, err
	}
	return res, nil
}

func (s *TaskService) FindLatestUnseenResults(limit uint32, page uint32) ([]*entity.TaskResultModel, error) {
	var res []*entity.TaskResultModel

	err := s.repo.GetByFieldSortedBatch(&res, "seen", false, "created_at", true, int(limit), int(page*limit))
	if err != nil {
		log.Debugf("[-] failed to find task results by taskid")
		return nil, err
	}
	return res, nil
}

func (s *TaskService) FindTaskResultOrError(resultId string) (*entity.TaskResultModel, error) {
	var taskRes entity.TaskResultModel
	err := s.repo.GetOneByID(&taskRes, resultId)

	if err != nil {
		return nil, err
	}
	return &taskRes, nil
}

func (s *TaskService) MarkTaskResultSeen(resultId string) error {
	taskRes, err := s.FindTaskResultOrError(resultId)
	if err != nil {
		return err
	}

	taskRes.Seen = true
	return s.repo.Save(taskRes)
}
