// internal/services/product.go
package services

import (
	"context"

	"github.com/google/uuid"

	"digital-market/internal/models"
	"digital-market/internal/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) error {
	return s.productRepo.Create(ctx, product)
}

func (s *ProductService) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, product *models.Product) error {
	return s.productRepo.Update(ctx, product)
}

func (s *ProductService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.productRepo.Delete(ctx, id)
}

func (s *ProductService) List(ctx context.Context, offset, limit int, categoryID *uuid.UUID) ([]models.Product, error) {
	return s.productRepo.List(ctx, offset, limit, categoryID)
}
