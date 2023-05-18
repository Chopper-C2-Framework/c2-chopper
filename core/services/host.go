package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type HostService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewHostService(db *orm.ORMConnection) *HostService {
	return &HostService{}
}
