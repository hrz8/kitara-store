package repository

import (
	"errors"

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

	ordersProducts := make([]models.OrdersProducts, 0)
	var priceAmount uint64 = 0
	for _, productPayload := range o.Products {
		product, err := h.productRepository.GetByID(db, productPayload.ID)
		if err != nil {
			return nil, err
		}

		// qtyAvailabe := product.QtyTotal - product.QtyReserved
		if product.QtyTotal < productPayload.Qty {
			return nil, errors.New("qty is bigger than available stock")
		}

		priceAmount += product.Price * productPayload.Qty
		opid, _ := uuid.NewV4()
		ordersProducts = append(ordersProducts, models.OrdersProducts{
			ID:        opid,
			ProductID: product.ID,
			Qty:       productPayload.Qty,
		})

		h.productRepository.UpdateOne(db, product, &models.Product{
			QtyTotal: product.QtyTotal - productPayload.Qty,
			// QtyReserved: product.QtyReserved + productPayload.Qty,
		})
	}

	id, _ := uuid.NewV4()
	order, err := h.repository.Create(db, &models.Order{
		ID:          id,
		PriceAmount: priceAmount,
		Status:      "RESERVED",
		Products:    ordersProducts,
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}
