package repository

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	ProductDomain "github.com/hrz8/kitara-store/domains/product"
	"github.com/hrz8/kitara-store/models"
)

type (
	handler struct {
		db *gorm.DB
	}
)

// NewRepository return implementation of methods in transaction.Repositoru
func NewRepository(db *gorm.DB) ProductDomain.Repository {
	return &handler{
		db: db,
	}
}

func (h handler) Create(db *gorm.DB, p *models.Product) (*models.Product, error) {
	if err := db.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (h handler) GetByID(db *gorm.DB, id uuid.UUID) (*models.Product, error) {
	p := models.Product{}
	if err := db.Take(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (h handler) UpdateOne(db *gorm.DB, p *models.Product, np *models.Product) (*models.Product, error) {
	if err := db.Model(&p).Updates(np).Error; err != nil {
		return nil, err
	}
	return np, nil
}
