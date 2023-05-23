package services

import "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

type Services struct {
	TeamService   ITeamService
	UserService   IUserService
	AgentService  IAgentService
	HostService   IHostService
	TaskService   ITaskService
	ReportService IReportService
}

type ITeamService interface {
	GetAll() ([]entity.TeamModel, error)
	CreateTeam(newTeam *entity.TeamModel) error
	AddMemberToTeam(teamId string, userId string) error
	DeleteTeam(teamId string) error
	FindOne(id string) (*entity.TeamModel, error)
	UpdateTeam(toUpdateTeamId string, toUpdateTeam *entity.TeamModel) (*entity.TeamModel, error)
}

type IUserService interface {
	CreateUser(newUser *entity.UserModel) error
	UpdateUser(id string, updatedUser *entity.UserModel) error
	UpdateUserPassword(id string, newPassword string) error
	FindUserOrError(id string) (*entity.UserModel, error)
	FindAll() ([]entity.UserModel, error)
}

type IAgentService interface {
	CreateAgent(agent *entity.AgentModel) error
	FindAgentOrError(id string) (*entity.AgentModel, error)
	ConnectAgent(id string) (*entity.AgentModel, error)

	UpdateAgent(agent *entity.AgentModel) error
}

type IHostService interface{}

type ITaskService interface {
	CreateTask(task *entity.TaskModel) error
	DeleteTask(task *entity.TaskModel) error

	FindTaskForAgent(agentId string) (*entity.TaskModel, error)
	FindTaskOrError(taskId string) (*entity.TaskModel, error)
}

type IReportService interface{}
