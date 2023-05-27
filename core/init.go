package core

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	services "github.com/chopper-c2-framework/c2-chopper/core/services"
)

func InitServices(db *orm.ORMConnection) services.Services {

	userService := services.NewUserService(db)
	return services.Services{
		TeamService:   services.NewTeamService(db),
		AgentService:  services.NewAgentService(db),
		HostService:   services.NewHostService(db),
		TaskService:   services.NewTaskService(db),
		ReportService: services.NewReportService(db),
		AuthService:   services.NewAuthService(userService),
	}
}
