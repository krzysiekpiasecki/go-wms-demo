package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.ProductRepository = (*ProductRepository)(nil)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	return r.db.QueryRow(
		context.Background(),
		`
        INSERT INTO products (name)
        VALUES ($1)
        RETURNING id, created_at
        `,
		product.Name,
	).Scan(
		&product.ID,
		&product.CreatedAt,
	)
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

func (r *ProductRepository) List() ([]domain.Product, error) {
	rows, err := r.db.Query(
		context.Background(),
		`

SELECT
    p.id,
    p.name,
    COALESCE(i.quantity, 0),
    p.created_at
FROM products p
LEFT JOIN inventory i
    ON i.product_id = p.id
ORDER BY p.id

        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Quantity,
			&product.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(
			products,
			product,
		)
	}

	return products, nil
}
