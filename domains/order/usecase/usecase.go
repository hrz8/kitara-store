package repository

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	OrderDomain "github.com/hrz8/kitara-store/domains/order"
	ProductDomain "github.com/hrz8/kitara-store/domains/product"
	"github.com/hrz8/kitara-store/models"
	"github.com/hrz8/kitara-store/shared/lib"
)

type (
	handler struct {
		repository        OrderDomain.Repository
		productRepository ProductDomain.Repository
	}
)

// NewUsecase return implementation of methods in order-domain.Repository
func NewUsecase(repo OrderDomain.Repository, productRepo ProductDomain.Repository) OrderDomain.Usecase {
	return &handler{
		repository:        repo,
		productRepository: productRepo,
	}
}

func (h handler) Create(c echo.Context, o *models.CreateOrderPayload) (*models.Order, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession
	trx := db.Begin()

	ordersProducts := make([]models.OrdersProducts, 0)
	var priceAmount uint64 = 0
	for _, productPayload := range o.Products {
		newProduct, err := h.productRepository.UpdateStock(trx, productPayload.ID, productPayload.Qty)
		if err != nil {
			trx.Rollback()
			if strings.Contains(err.Error(), "value is out of range") {
				return nil, errors.New("qty is bigger than available stock")
			}
			return nil, err
		}

		if newProduct.QtyTotal < 0 {
			trx.Rollback()
			return nil, errors.New("qty is bigger than available stock")
		}

		priceAmount += newProduct.Price * productPayload.Qty
		opid, _ := uuid.NewV4()
		ordersProducts = append(ordersProducts, models.OrdersProducts{
			ID:        opid,
			ProductID: newProduct.ID,
			Qty:       productPayload.Qty,
			Product:   *newProduct,
		})
	}

	id, _ := uuid.NewV4()
	order, err := h.repository.Create(trx, &models.Order{
		ID:          id,
		PriceAmount: priceAmount,
		Products:    ordersProducts,
	})
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return order, nil
}
