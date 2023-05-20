package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type TaskService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewTaskService(db *orm.ORMConnection) *TaskService {
	return &TaskService{}
}
