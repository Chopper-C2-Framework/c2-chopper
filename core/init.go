package core

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	services "github.com/chopper-c2-framework/c2-chopper/core/services"
)

func InitServices(db *orm.ORMConnection) services.Services {

	return services.Services{
		TeamService:     services.NewTeamService(db),
		UserService:     services.NewUserService(db),
		ListenerService: services.NewListenerService(db),
		HostService:     services.NewHostService(db),
		TaskService:     services.NewTaskService(db),
		ReportService:   services.NewReportService(db),
	}
}
