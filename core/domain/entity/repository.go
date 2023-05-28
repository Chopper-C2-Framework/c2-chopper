package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type gormRepository struct {
	db           *gorm.DB
	logger       Logger
	defaultJoins []string
}

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound = gorm.ErrRecordNotFound
)

// NewGormRepository returns a new base repository that implements TransactionRepository
func NewGormRepository(db *gorm.DB, logger Logger, defaultJoins ...string) TransactionRepository {
	return &gormRepository{
		defaultJoins: defaultJoins,
		logger:       logger,
		db:           db,
	}
}

func (r *gormRepository) DB() *gorm.DB {
	return r.DBWithPreloads(nil)
}

func (r *gormRepository) GetAll(target interface{}, preloads ...string) error {
	r.logger.Debugf("Executing GetAll on %T", target)

	res := r.DBWithPreloads(preloads).Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetBatch(target interface{}, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetBatch on %T", target)

	res := r.DBWithPreloads(preloads).
		Unscoped().
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetSortedBatch(target interface{}, column string, ascending bool, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetSortedBatch on %T", target)

	order := column
	if !ascending {
		order += " DESC"
	}

	res := r.DBWithPreloads(preloads).
		Unscoped().
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetByFieldSortedBatch(target interface{}, field string, value interface{}, sortColumn string, ascending bool, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetByFieldSortedBatch on %T with %v = %v", target, field, value)

	order := sortColumn
	if !ascending {
		order += " DESC"
	}

	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%v = ?", field), value).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetWhere(target interface{}, condition string, preloads ...string) error {
	r.logger.Debugf("Executing GetWhere on %T with %v ", target, condition)

	res := r.DBWithPreloads(preloads).
		Where(condition).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetWhereBatch(target interface{}, condition string, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetWhere on %T with %v ", target, condition)

	res := r.DBWithPreloads(preloads).
		Where(condition).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetByField(target interface{}, field string, value interface{}, preloads ...string) error {
	r.logger.Debugf("Executing GetByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%v = ?", field), value).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetByFields(target interface{}, filters map[string]interface{}, preloads ...string) error {
	r.logger.Debugf("Executing GetByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetByFieldBatch(target interface{}, field string, value interface{}, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%v = ?", field), value).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetByFieldsBatch(target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) error {
	r.logger.Debugf("Executing GetByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) GetOneByField(target interface{}, field string, value interface{}, preloads ...string) error {
	r.logger.Debugf("Executing GetOneByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%v = ?", field), value).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) GetOneByFields(target interface{}, filters map[string]interface{}, preloads ...string) error {
	r.logger.Debugf("Executing FindOneByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.First(target)
	return r.HandleOneError(res)
}

func (r *gormRepository) GetOneByID(target interface{}, id string, preloads ...string) error {
	r.logger.Debugf("Executing GetOneByID on %T with ID %v", target, id)

	res := r.DBWithPreloads(preloads).
		Where("id = ?", id).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) Create(target interface{}) error {
	r.logger.Debugf("Executing Create on %T", target)

	res := r.db.Create(target)
	return r.HandleError(res)
}

func (r *gormRepository) CreateTx(target interface{}, tx *gorm.DB) error {
	r.logger.Debugf("Executing Create on %T", target)

	res := tx.Create(target)
	return r.HandleError(res)
}

func (r *gormRepository) Save(target interface{}) error {
	r.logger.Debugf("Executing Save on %T", target)

	res := r.db.Save(target)
	return r.HandleError(res)
}

func (r *gormRepository) Update(target interface{}, updates interface{}) error {
	r.logger.Debugf("Executing Save on %T", target)

	res := r.db.Model(target).Updates(updates)
	return r.HandleError(res)
}

func (r *gormRepository) SaveTx(target interface{}, tx *gorm.DB) error {
	r.logger.Debugf("Executing Save on %T", target)

	res := tx.Save(target)
	return r.HandleError(res)
}

func (r *gormRepository) Delete(target interface{}) error {
	r.logger.Debugf("Executing Delete on %T", target)

	res := r.db.Delete(target)
	return r.HandleError(res)
}

func (r *gormRepository) DeleteTx(target interface{}, tx *gorm.DB) error {
	r.logger.Debugf("Executing Delete on %T", target)

	res := tx.Delete(target)
	return r.HandleError(res)
}

func (r *gormRepository) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("error: %w", res.Error)
		r.logger.Error(err)
		return err
	}

	return nil
}

func (r *gormRepository) HandleOneError(res *gorm.DB) error {
	if err := r.HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return ErrNotFound
	}

	return nil
}

func (r *gormRepository) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := r.db

	for _, join := range r.defaultJoins {
		dbConn = dbConn.Joins(join)
	}

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}
