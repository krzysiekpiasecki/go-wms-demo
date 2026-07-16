package service

import (
	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

type OrderService struct {
	orderRepository     repository.OrderRepository
	inventoryRepository repository.InventoryRepository
	orderItemRepository repository.OrderItemRepository
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	orderItemRepository repository.OrderItemRepository,
	inventoryRepository repository.InventoryRepository,
) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
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

func (s *OrderService) CreateOrder(productID int64, quantity int, comment *string) error {
	err := s.CanFulfillOrder(productID, quantity)
	if err != nil {
		return err
	}

	order := &domain.Order{
		Status:  domain.OrderStatusNew,
		Comment: comment,
	}

	err = s.orderRepository.Create(order)
	if err != nil {
		return err
	}

	item := &domain.OrderItem{
		OrderID:   order.ID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.orderItemRepository.Create(item)
}

func (s *OrderService) GetOrder(id int64) (*domain.Order, error) {
	order, err := s.orderRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	items, err := s.orderItemRepository.GetByOrderID(id)
	if err != nil {
		return nil, err
	}

	order.Items = items

	return order, nil
}

func (s *OrderService) UpdateStatus(id int64, status string) error {
	return s.orderRepository.UpdateStatus(id, status)
}
