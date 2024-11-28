package controllers

import (
	"net/http"
	"strconv"

	"github.com/gareisdev/bar-control/internal/core/entities"
	"github.com/gareisdev/bar-control/internal/core/usecases"
	"github.com/labstack/echo/v4"
)

type MenuController struct {
	menuUsecase *usecases.MenuUsecase
}

func NewMenuController(menuUsecase *usecases.MenuUsecase) *MenuController {
	return &MenuController{menuUsecase: menuUsecase}
}

func (mc *MenuController) GetAll(c echo.Context) error {
	items, err := mc.menuUsecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, items)
}

func (mc *MenuController) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	item, err := mc.menuUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if item == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Menu item not found"})
	}
	return c.JSON(http.StatusOK, item)
}

func (mc *MenuController) Create(c echo.Context) error {
	var menuItem entities.MenuItem
	if err := c.Bind(&menuItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	id, err := mc.menuUsecase.Create(c.Request().Context(), &menuItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]int64{"id": id})
}

func (mc *MenuController) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var menuItem entities.MenuItem
	if err := c.Bind(&menuItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	menuItem.ID = id
	if err := mc.menuUsecase.Update(c.Request().Context(), &menuItem); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Menu item updated successfully"})
}

func (mc *MenuController) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := mc.menuUsecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Menu item deleted successfully"})
}
