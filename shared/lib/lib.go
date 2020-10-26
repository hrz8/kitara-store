package lib

import (
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ProductRepository "github.com/hrz8/kitara-store/domains/product/repository"
	"github.com/hrz8/kitara-store/models"
)

type (
	// AppContext return custom application context
	AppContext struct {
		echo.Context
		MysqlSession *gorm.DB
	}

	// CustomValidator return custom request validator
	CustomValidator struct {
		Validator *validator.Validate
	}
)

// Validate will validate given input with related struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// Migrate will doing migration for DB
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.Order{},
		&models.OrdersProducts{},
	)

	migrationRepo := ProductRepository.NewRepository(db)
	migrationRepo.Create(db, &models.Product{
		ID:          uuid.FromStringOrNil("35731de0-e646-4379-a0f1-b69e74742e0a"),
		Name:        "Prod Nomer Satu",
		Price:       13500,
		QtyTotal:    17,
		QtyReserved: 0,
	})
	migrationRepo.Create(db, &models.Product{
		ID:          uuid.FromStringOrNil("6da51a6f-10c0-404c-82f1-2dce60d720a4"),
		Name:        "Prod Nomer Dua",
		Price:       6000,
		QtyTotal:    28,
		QtyReserved: 0,
	})
	migrationRepo.Create(db, &models.Product{
		ID:          uuid.FromStringOrNil("0bde6df2-f505-401f-882c-808855c2871d"),
		Name:        "Prod Nomer Tg",
		Price:       17000,
		QtyTotal:    11,
		QtyReserved: 0,
	})
}
