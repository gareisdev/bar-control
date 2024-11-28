package usecases

import (
	"context"
	"errors"

	"github.com/gareisdev/bar-control/internal/core/entities"
	"github.com/gareisdev/bar-control/internal/core/interfaces"
)

type MenuUsecase struct {
	menuRepo interfaces.MenuRepository
}

func NewMenuUsecase(menuRepo interfaces.MenuRepository) *MenuUsecase {
	return &MenuUsecase{menuRepo: menuRepo}
}

func (uc *MenuUsecase) GetAll(ctx context.Context) ([]entities.MenuItem, error) {
	return uc.menuRepo.GetAll(ctx)
}

func (uc *MenuUsecase) GetByID(ctx context.Context, id int64) (*entities.MenuItem, error) {
	menuItem, err := uc.menuRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if menuItem == nil {
		return nil, errors.New("menu item not found")
	}
	return menuItem, nil
}

func (uc *MenuUsecase) Create(ctx context.Context, menuItem *entities.MenuItem) (int64, error) {
	if menuItem.Name == "" {
		return 0, errors.New("menu item name is required")
	}
	return uc.menuRepo.Create(ctx, menuItem)
}

func (uc *MenuUsecase) Update(ctx context.Context, menuItem *entities.MenuItem) error {
	if menuItem.ID == 0 {
		return errors.New("menu item ID is required")
	}
	return uc.menuRepo.Update(ctx, menuItem)
}

func (uc *MenuUsecase) Delete(ctx context.Context, id int64) error {
	return uc.menuRepo.Delete(ctx, id)
}
