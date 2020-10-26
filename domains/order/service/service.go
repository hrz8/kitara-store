package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	OrderDomain "github.com/hrz8/kitara-store/domains/order"
	"github.com/hrz8/kitara-store/models"
)

type (
	handler struct {
		usecase OrderDomain.Usecase
	}
)

// InitService will return REST of player-domain
func InitService(e *echo.Echo, usecase OrderDomain.Usecase) {
	h := handler{
		usecase: usecase,
	}

	e.POST("/api/v1/orders", h.Create)
}

func (h handler) Create(c echo.Context) error {
	payload := &models.CreateOrderPayload{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	order, err := h.usecase.Create(c, payload)
	if err != nil {
		var status int
		switch err.Error() {
		case "qty is bigger than available stock":
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, echo.Map{
			"status": status,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   order,
	})
}
