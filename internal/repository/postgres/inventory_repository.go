package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.InventoryRepository = (*InventoryRepository)(nil)

type InventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) GetByProductID(productID int64) (*domain.Inventory, error) {
	var inventory domain.Inventory

	err := r.db.QueryRow(
		context.Background(),
		`
        SELECT product_id, quantity
        FROM inventory
        WHERE product_id = $1
        `,
		productID,
	).Scan(
		&inventory.ProductID,
		&inventory.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepository) UpdateQuantity(productID int64, quantity int) error {
	_, err := r.db.Exec(
		context.Background(),
		`
        UPDATE inventory
        SET quantity = $1
        WHERE product_id = $2
        `,
		quantity,
		productID,
	)

	return err
}

func (r *InventoryRepository) Create(
	productID int64,
	quantity int,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`
        INSERT INTO inventory (
            product_id,
            quantity
        )
        VALUES ($1, $2)
        `,
		productID,
		quantity,
	)

	return err
}
func (r *InventoryRepository) DecreaseStock(
	productID int64,
	quantity int,
) error {
	_, err := r.db.Exec(
		context.Background(),
		`
        UPDATE inventory
        SET quantity = quantity - $1
        WHERE product_id = $2
        `,
		quantity,
		productID,
	)

	return err
}
