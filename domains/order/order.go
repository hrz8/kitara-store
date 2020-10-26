package container

import (
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"

	"github.com/hrz8/kitara-store/models"
)

type (
	// Repository is an interface of Order domain for user model method
	Repository interface {
		Create(db *gorm.DB, o *models.Order) (*models.Order, error)
	}

	// Usecase is an interface of Order domain for player sharable method
	Usecase interface {
		Create(c echo.Context, o *models.CreateOrderPayload) (*models.Order, error)
	}
)
