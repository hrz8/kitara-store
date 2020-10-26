package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	OrderRepository "github.com/hrz8/kitara-store/domains/order/repository"
	OrderService "github.com/hrz8/kitara-store/domains/order/service"
	OrderUsecase "github.com/hrz8/kitara-store/domains/order/usecase"

	ProductRepository "github.com/hrz8/kitara-store/domains/product/repository"
	// OrderUsecase "github.com/hrz8/kitara-store/domains/order/usecase"

	"github.com/hrz8/kitara-store/shared/config"
	"github.com/hrz8/kitara-store/shared/container"
	"github.com/hrz8/kitara-store/shared/database"
	"github.com/hrz8/kitara-store/shared/lib"
)

func main() {
	e := echo.New()

	appContainer := container.DefaultContainer()
	appConfig := appContainer.MustGet("shared.config").(config.AppConfigInterface)
	mysql := appContainer.MustGet("shared.mysql").(database.MysqlInterface)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	mysqlSess, err := mysql.Connect()
	if err != nil {
		panic(fmt.Sprintf("[ERROR] failed open mysql connection: %s", err.Error()))
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "migrate" {
			lib.Migrate(mysqlSess)
		}
	}

	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &lib.AppContext{
				Context:      c,
				MysqlSession: mysqlSess,
			}
			return next(ac)
		}
	})

	orderRepo := OrderRepository.NewRepository(mysqlSess)
	productRepo := ProductRepository.NewRepository(mysqlSess)

	orderUsecase := OrderUsecase.NewUsecase(orderRepo, productRepo)

	OrderService.InitService(e, orderUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.GetAppPort())))
}
