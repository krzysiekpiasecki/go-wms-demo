package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(order *domain.Order) error {
	return r.db.QueryRow(
		context.Background(),
		`
		INSERT INTO orders(status, comment)
		VALUES ($1, $2)
		RETURNING id
        `,
		order.Status,
		order.Comment,
	).Scan(&order.ID)
}

func (r *OrderRepository) GetByID(id int64) (*domain.Order, error) {
	var order domain.Order

	err := r.db.QueryRow(
		context.Background(),
		`
        SELECT id, status, comment, created_at
        FROM orders
        WHERE id = $1
        `,
		id,
	).Scan(
		&order.ID,
		&order.Status,
		&order.Comment,
		&order.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrOrderNotFound
		}

		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdateStatus(id int64, status string) error {
	result, err := r.db.Exec(
		context.Background(),
		`
        UPDATE orders
        SET status = $1
        WHERE id = $2
        `,
		status,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrOrderNotFound
	}

	return nil
}

func (r *OrderRepository) List() ([]domain.Order, error) {
	rows, err := r.db.Query(
		context.Background(),
		`
        SELECT
            id,
            status,
            comment,
            created_at
        FROM orders
        ORDER BY id DESC
        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		var order domain.Order

		err := rows.Scan(
			&order.ID,
			&order.Status,
			&order.Comment,
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(
			orders,
			order,
		)
	}

	return orders, nil
}
