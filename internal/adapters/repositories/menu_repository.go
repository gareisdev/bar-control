package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gareisdev/bar-control/internal/core/entities"
	"github.com/jmoiron/sqlx"
)

type MenuRepository interface {
	GetAll(ctx context.Context) ([]entities.MenuItem, error)
	GetByID(ctx context.Context, id int64) (*entities.MenuItem, error)
	Create(ctx context.Context, menuItem *entities.MenuItem) (int64, error)
	Update(ctx context.Context, menuItem *entities.MenuItem) error
	Delete(ctx context.Context, id int64) error
}

type menuRepository struct {
	db *sqlx.DB
}

// Construct
func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r *menuRepository) GetAll(ctx context.Context) ([]entities.MenuItem, error) {
	var menuItems []entities.MenuItem
	query := `SELECT id, name, description, price, available, created_at FROM menu_items`

	err := r.db.SelectContext(ctx, &menuItems, query)
	if err != nil {
		return nil, err
	}

	return menuItems, nil
}

func (r *menuRepository) GetByID(ctx context.Context, id int64) (*entities.MenuItem, error) {
	var menuItem entities.MenuItem
	query := `SELECT id, name, description, price, available, created_at FROM menu_items WHERE id = $1`

	err := r.db.GetContext(ctx, &menuItem, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // No encontrado
	} else if err != nil {
		return nil, err
	}

	return &menuItem, nil
}

func (r *menuRepository) Create(ctx context.Context, menuItem *entities.MenuItem) (int64, error) {
	query := `INSERT INTO menu_items (name, description, price, available) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	err := r.db.QueryRowContext(ctx, query, menuItem.Name, menuItem.Description, menuItem.Price, menuItem.Available).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *menuRepository) Update(ctx context.Context, menuItem *entities.MenuItem) error {
	query := `UPDATE menu_items SET name = $1, description = $2, price = $3, available = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query, menuItem.Name, menuItem.Description, menuItem.Price, menuItem.Available, menuItem.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *menuRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM menu_items WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
