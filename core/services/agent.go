package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type AgentService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewAgentService(db *orm.ORMConnection) *AgentService {
	return &AgentService{}
}
