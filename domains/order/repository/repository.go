package repository

import (
	"gorm.io/gorm"

	OrderDomain "github.com/hrz8/kitara-store/domains/order"
	"github.com/hrz8/kitara-store/models"
)

type (
	handler struct {
		db *gorm.DB
	}
)

// NewRepository return implementation of methods in transaction.Repositoru
func NewRepository(db *gorm.DB) OrderDomain.Repository {
	return &handler{
		db: db,
	}
}

func (h handler) Create(db *gorm.DB, o *models.Order) (*models.Order, error) {
	if err := db.Create(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}
