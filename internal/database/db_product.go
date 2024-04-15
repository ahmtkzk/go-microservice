package database

import (
	"context"
	"go-microservice/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	*DatabaseRepository[model.Products]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		&DatabaseRepository[model.Products]{db: db},
	}
}

func (pr *ProductRepository) GetAllProducts(ctx context.Context) []model.Products {
	result, _ := pr.GetAll(ctx)
	return result
}

func (pr *ProductRepository) GetProductByName(ctx context.Context, name string) model.Products {
	result, _ := pr.GetByParam(ctx, &model.Products{Name: name})
	return result
}
