package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.ProductRepository = (*ProductRepository)(nil)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	_, err := r.db.Exec(
		context.Background(),
		`
        INSERT INTO products(name)
        VALUES ($1)
        `,
		product.Name,
	)

	return err
}

func (r *ProductRepository) GetByID(id int64) (*domain.Product, error) {
	var product domain.Product

	err := r.db.QueryRow(
		context.Background(),
		`
        SELECT id, name, created_at
        FROM products
        WHERE id = $1
        `,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrProductNotFound
		}

		return nil, err
	}

	return &product, nil
}
