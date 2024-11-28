package entities

import "time"

type MenuItem struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Available   bool      `db:"available"`
	CreatedAt   time.Time `db:"created_at"`
}
