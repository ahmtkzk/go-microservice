package database

import (
	"context"
	"fmt"
	"go-microservice/internal/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DbOperations interface {
	Ready() bool
	GetAll(ctx context.Context) ([]any, error)
	GetByParam(ctx context.Context, param any) (any, error)
	Create(ctx context.Context, data any) error
	Update(ctx context.Context, data any) error
	Delete(ctx context.Context, model any, param any) error
}

type DbConnection struct {
	DB *gorm.DB
}

func NewDatabaseConnection() DbConnection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		constants.Host,
		constants.User,
		constants.Password,
		constants.Dbname,
		constants.Port,
		constants.SSLMode)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true})

	return DbConnection{DB: db}
}

func (db *DbConnection) Ready() bool {
	var ready string
	tx := db.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (db *DbConnection) GetAll(ctx context.Context) ([]any, error) {
	var list []any
	result := db.DB.WithContext(ctx).Find(&list)
	return list, result.Error
}

func (db *DbConnection) GetByParam(ctx context.Context, param any) (any, error) {
	var model any
	result := db.DB.WithContext(ctx).Where(param).Find(&model)
	return model, result.Error
}

func (db *DbConnection) Create(ctx context.Context, data any) error {
	result := db.DB.WithContext(ctx).Create(data)
	return result.Error
}

func (db *DbConnection) Update(ctx context.Context, data any) error {
	result := db.DB.WithContext(ctx).Model(data)
	return result.Error
}

func (db *DbConnection) Delete(ctx context.Context, model any, param any) error {
	result := db.DB.WithContext(ctx).Delete(model, param)
	return result.Error
}
