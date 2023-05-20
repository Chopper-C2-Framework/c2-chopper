package services

import "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

type Services struct {
	TeamService   ITeamService
	AuthService   IAuthService
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
	FindUserByIdOrError(id string) (*entity.UserModel, error)
	FindUserByUsernameOrError(username string) (*entity.UserModel, error)
	FindAll() ([]entity.UserModel, error)
}

type IAuthService interface {
	Login(username string, password string) (string, error)
	Register(username string, password string) (string, error)
}

type IAgentService interface{}

type IHostService interface{}

type ITaskService interface{}

type IReportService interface{}
