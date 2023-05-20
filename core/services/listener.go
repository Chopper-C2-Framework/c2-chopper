package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type ListenerService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewListenerService(db *orm.ORMConnection) *ListenerService {
	return &ListenerService{}
}
