package interfaces

import (
	"context"

	"github.com/gareisdev/bar-control/internal/core/entities"
)

// MenuRepository es la interfaz para las operaciones de `menu_items`.
type MenuRepository interface {
	GetAll(ctx context.Context) ([]entities.MenuItem, error)
	GetByID(ctx context.Context, id int64) (*entities.MenuItem, error)
	Create(ctx context.Context, menuItem *entities.MenuItem) (int64, error)
	Update(ctx context.Context, menuItem *entities.MenuItem) error
	Delete(ctx context.Context, id int64) error
}
