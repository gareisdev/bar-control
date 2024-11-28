package http

import (
	"github.com/gareisdev/bar-control/internal/adapters/controllers"
	"github.com/gareisdev/bar-control/internal/core/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupServer(menuUsecase *usecases.MenuUsecase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	menuController := controllers.NewMenuController(menuUsecase)

	menu := e.Group("/menu-items")
	menu.GET("", menuController.GetAll)
	menu.GET("/:id", menuController.GetByID)
	menu.POST("", menuController.Create)
	menu.PUT("/:id", menuController.Update)
	menu.DELETE("/:id", menuController.Delete)

	return e
}
