// internal/repositories/product.go
package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"digital-market/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, "id = ?", id).Error
}

func (r *ProductRepository) List(ctx context.Context, offset, limit int, categoryID *uuid.UUID) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).Offset(offset).Limit(limit)
	
	if categoryID != nil {
		query = query.Where("category_id = ?", categoryID)
	}
	
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	
	return products, nil
}
