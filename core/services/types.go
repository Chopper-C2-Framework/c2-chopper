package services

import entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

type ITeamService interface {
	CreateTeam(newTeam *entity.TeamModel) error
	AddMemberToTeam(teamId string, user_id string) error
	UpdateTeam(teamId string, newTeam *entity.TeamModel) error
	DeleteTeam(teamId string) error
}

type IUserService interface {
	CreateUser(newUser *entity.UserModel) error
	UpdateUser(id string, updatedUser *entity.UserModel) error
	UpdateUserPassword(id string, newPassword string) error
	FindUserOrError(id string) (*entity.UserModel, error)
	FindAll() ([]entity.UserModel, error)
}

type IListenerService interface{}

type IHostService interface{}

type ITaskService interface{}

type IReportService interface{}
