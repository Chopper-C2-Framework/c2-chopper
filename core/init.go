package core

import (
	"github.com/chopper-c2-framework/c2-chopper/core/config"
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	services "github.com/chopper-c2-framework/c2-chopper/core/services"
)

func InitServices(db *orm.ORMConnection, frameworkConfig config.Config) services.Services {

	userService := services.NewUserService(db)
	return services.Services{
		TeamService:         services.NewTeamService(db),
		AgentService:        services.NewAgentService(db),
		HostService:         services.NewHostService(db),
		TaskService:         services.NewTaskService(db),
		ReportService:       services.NewReportService(db),
		PluginResultService: services.NewPluginResultService(db),
		AuthService:         services.NewAuthService(userService, frameworkConfig),
	}
}
