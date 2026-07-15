package service

import (
	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

type OrderService struct {
	orderRepository     repository.OrderRepository
	inventoryRepository repository.InventoryRepository
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	inventoryRepository repository.InventoryRepository,
) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *OrderService) HasEnoughStock(productID int64, quantity int) (bool, error) {

	if quantity <= 0 {
		return false, domain.ErrInvalidQuantity
	}

	inventory, err := s.inventoryRepository.GetByProductID(productID)
	if err != nil {
		return false, err
	}

	return inventory.Quantity >= quantity, nil
}

func (s *OrderService) CanFulfillOrder(productID int64, quantity int) error {
	hasEnoughStock, err := s.HasEnoughStock(productID, quantity)
	if err != nil {
		return err
	}

	if !hasEnoughStock {
		return domain.ErrInsufficientStock
	}

	return nil
}

func (s *OrderService) CreateOrder(productID int64, quantity int) error {
	err := s.CanFulfillOrder(productID, quantity)
	if err != nil {
		return err
	}

	order := &domain.Order{
		Status: domain.OrderStatusNew,
	}

	return s.orderRepository.Create(order)
}
