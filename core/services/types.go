package services

import "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

type Services struct {
	TeamService         ITeamService
	AuthService         IAuthService
	AgentService        IAgentService
	HostService         IHostService
	TaskService         ITaskService
	ReportService       IReportService
	PluginResultService IPluginResultService
}

type IPluginResultService interface {
	CreatePluginResult(newRes *entity.PluginResultModel) error
	GetPluginResultsOrError(path string) ([]*entity.PluginResultModel, error)
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
	FindUserByIdOrError(id string) (*entity.UserModel, error)
	FindUserByUsernameOrError(username string) (*entity.UserModel, error)
	FindAll() ([]entity.UserModel, error)
}

type IAuthService interface {
	Login(username string, password string) (string, error)
	Register(username string, password string) (string, error)
	ParseToken(token string) (*JWTData, error)
}

type IAgentService interface {
	CreateAgent(agent *entity.AgentModel) error
	FindAgentOrError(id string) (*entity.AgentModel, error)
	FindAllAgents() ([]*entity.AgentModel, error)
	ConnectAgent(id string, agentInfo *entity.AgentModel) (*entity.AgentModel, error)

	UpdateAgent(target *entity.AgentModel, updates *entity.AgentModel) error
	SaveAgent(agent *entity.AgentModel) error
}

type IHostService interface{}

type ITaskService interface {
	CreateTask(task *entity.TaskModel) error
	DeleteTask(task *entity.TaskModel) error

	FindAllTasks() ([]*entity.TaskModel, error)

	FindTasksForAgent(agentId string) ([]*entity.TaskModel, error)
	FindUnexecutedTasksForAgent(agentId string) ([]*entity.TaskModel, error)
	FindTaskOrError(taskId string) (*entity.TaskModel, error)

	CreateTaskResult(taskResult *entity.TaskResultModel) error
	FindTaskResults(taskId string) ([]*entity.TaskResultModel, error)
	FindTaskResultOrError(resultId string) (*entity.TaskResultModel, error)
	MarkTaskResultSeen(resultId string) error

	FindLatestResults(limit uint32, page uint32) ([]*entity.TaskResultModel, error)
}

type IReportService interface{}
