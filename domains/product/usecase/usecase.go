package repository

import (
	"github.com/labstack/echo/v4"

	ProductDomain "github.com/hrz8/kitara-store/domains/product"
	"github.com/hrz8/kitara-store/models"
	"github.com/hrz8/kitara-store/shared/lib"
)

type (
	handler struct {
		repository ProductDomain.Repository
	}
)

// NewUsecase return implementation of methods in product-domain.Repository
func NewUsecase(repo ProductDomain.Repository) ProductDomain.Usecase {
	return &handler{
		repository: repo,
	}
}

func (h handler) Create(c echo.Context, p *models.Product) (*models.Product, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	product, err := h.repository.Create(db, p)
	if err != nil {
		return nil, err
	}

	return product, nil
}
