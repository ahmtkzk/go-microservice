package database

import (
	"context"
	"gorm.io/gorm"
)

type DatabaseRepository[T any] struct {
	db *gorm.DB
}

func (dbr *DatabaseRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	var models []T
	result := dbr.db.WithContext(ctx).Find(&models)
	return models, result.Error
}

func (dbr *DatabaseRepository[T]) GetByParam(ctx context.Context, param any) (T, error) {
	var model T
	result := dbr.db.WithContext(ctx).Where(param).Find(&model)
	return model, result.Error
}

func (dbr *DatabaseRepository[T]) Create(ctx context.Context, data T) error {
	result := dbr.db.WithContext(ctx).Create(data)
	return result.Error
}

func (dbr *DatabaseRepository[T]) Update(ctx context.Context, data T) error {
	result := dbr.db.WithContext(ctx).Model(data)
	return result.Error
}

func (dbr *DatabaseRepository[T]) Delete(ctx context.Context, model T, param any) error {
	result := dbr.db.WithContext(ctx).Delete(model, param)
	return result.Error
}
