package repository

import "github.com/kpiasecki/wms/internal/domain"

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id int64) (*domain.Product, error)
}

type InventoryRepository interface {
	GetByProductID(productID int64) (*domain.Inventory, error)
	UpdateQuantity(productID int64, quantity int) error
}

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id int64) (*domain.Order, error)
	UpdateStatus(id int64, status string) error
}

type OrderItemRepository interface {
	Create(item *domain.OrderItem) error
	GetByOrderID(orderID int64) ([]domain.OrderItem, error)
}
