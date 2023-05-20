package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
)

type ReportService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewReportService(db *orm.ORMConnection) *ReportService {
	return &ReportService{}
}
