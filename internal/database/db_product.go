package database

import (
	"context"
	"go-microservice/internal/models"
)

type ProductOperations interface {
	GetAllProducts(ctx context.Context) ([]models.Products, error)
	GetProductById(ctx context.Context) (models.Products, error)
}

func (db *DbConnection) GetAllProducts(ctx context.Context) ([]models.Products, error) {
	return nil, nil
}

func (db *DbConnection) GetProductById(ctx context.Context) (models.Products, error) {
	a, _ := db.GetByParam(ctx, "asd")
	return a, nil
}
