package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.OrderItemRepository = (*OrderItemRepository)(nil)

type OrderItemRepository struct {
	db *pgx.Conn
}

func NewOrderItemRepository(db *pgx.Conn) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

func (r *OrderItemRepository) Create(item *domain.OrderItem) error {
	return r.db.QueryRow(
		context.Background(),
		`
        INSERT INTO order_items(order_id, product_id, quantity)
        VALUES ($1, $2, $3)
        RETURNING id
        `,
		item.OrderID,
		item.ProductID,
		item.Quantity,
	).Scan(&item.ID)
}

func (r *OrderItemRepository) GetByOrderID(orderID int64) ([]domain.OrderItem, error) {
	rows, err := r.db.Query(
		context.Background(),
		`
        SELECT id, order_id, product_id, quantity
        FROM order_items
        WHERE order_id = $1
        `,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.OrderItem

	for rows.Next() {
		var item domain.OrderItem

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
