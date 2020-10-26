package lib

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
