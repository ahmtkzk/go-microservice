package database

import (
	"fmt"
	"go-microservice/internal/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DbOperations interface {
	Ready() bool
}

type DbConnection struct {
	DB *gorm.DB
}

func NewDatabaseConnection() DbOperations {
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

func (db DbConnection) Ready() bool {
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
