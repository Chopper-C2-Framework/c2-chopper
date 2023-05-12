package services

import entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

type ITeamService interface {
	CreateTeam(newTeam *entity.TeamModel) error
	AddMemberToTeam(teamId string, user_id string) error
	UpdateTeam(teamId string, newTeam *entity.TeamModel) error
	DeleteTeam(teamId string) error
}

type IUserService interface{}

type IListenerService interface{}

type IHostService interface{}

type ITaskService interface{}

type IReportService interface{}
