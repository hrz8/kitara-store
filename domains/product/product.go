package container

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	"github.com/hrz8/kitara-store/models"
)

type (
	// Repository is an interface of Product domain for user model method
	Repository interface {
		Create(db *gorm.DB, p *models.Product) (*models.Product, error)
		GetByID(db *gorm.DB, id uuid.UUID) (*models.Product, error)
		UpdateOne(db *gorm.DB, p *models.Product, np *models.Product) (*models.Product, error)
		UpdateStock(db *gorm.DB, id uuid.UUID, qty uint64) (*models.Product, error)
	}

	// Usecase is an interface of Product domain for player sharable method
	Usecase interface {
		Create(c echo.Context, p *models.Product) (*models.Product, error)
	}
)
