package service

import (
	"errors"
	"testing"

	"github.com/kpiasecki/wms/internal/domain"
)

type MockOrderRepository struct {
	createErr  error
	created    bool
	savedOrder *domain.Order
}

type MockOrderItemRepository struct {
	createErr error
	items     []domain.OrderItem
}

func (m *MockOrderItemRepository) Create(item *domain.OrderItem) error {
	return m.createErr
}

func (m *MockOrderItemRepository) GetByOrderID(orderID int64) ([]domain.OrderItem, error) {
	return m.items, nil
}

func (m *MockOrderRepository) Create(order *domain.Order) error {
	m.created = true
	m.savedOrder = order
	return m.createErr
}

func (m *MockOrderRepository) GetByID(id int64) (*domain.Order, error) {
	return nil, nil
}

func (m *MockOrderRepository) UpdateStatus(id int64, status string) error {
	return nil
}

func TestHasEnoughStock(t *testing.T) {
	tests := []struct {
		name        string
		inventory   *domain.Inventory
		err         error
		quantity    int
		expected    bool
		expectedErr error
	}{
		{
			name: "enough stock",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  100,
			},
			quantity: 50,
			expected: true,
		},
		{
			name: "not enough stock",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  10,
			},
			quantity: 50,
			expected: false,
		},
		{
			name:        "repository error",
			err:         domain.ErrProductNotFound,
			quantity:    50,
			expectedErr: domain.ErrProductNotFound,
		},
		{
			name:        "invalid quantity",
			quantity:    0,
			expectedErr: domain.ErrInvalidQuantity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inventoryRepo := &MockInventoryRepository{
				inventory: test.inventory,
				err:       test.err,
			}

			orderItemRepo := &MockOrderItemRepository{}

			service := NewOrderService(
				&MockOrderRepository{},
				orderItemRepo,
				inventoryRepo,
			)

			result, err := service.HasEnoughStock(1, test.quantity)

			if test.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if test.expectedErr != nil {
				if !errors.Is(err, test.expectedErr) {
					t.Fatalf("expected %v, got %v", test.expectedErr, err)
				}
				return
			}

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestCanFulfillOrder(t *testing.T) {
	tests := []struct {
		name        string
		inventory   *domain.Inventory
		quantity    int
		err         error
		expectedErr error
	}{
		{
			name: "can fulfill",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  100,
			},
			quantity: 50,
		},
		{
			name: "insufficient stock",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  10,
			},
			quantity:    50,
			expectedErr: domain.ErrInsufficientStock,
		},
		{
			name:        "invalid quantity",
			quantity:    0,
			expectedErr: domain.ErrInvalidQuantity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inventoryRepo := &MockInventoryRepository{
				inventory: test.inventory,
				err:       test.err,
			}

			orderItemRepo := &MockOrderItemRepository{}

			service := NewOrderService(
				&MockOrderRepository{},
				orderItemRepo,
				inventoryRepo,
			)

			err := service.CanFulfillOrder(1, test.quantity)

			if test.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if test.expectedErr != nil && !errors.Is(err, test.expectedErr) {
				t.Fatalf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name        string
		inventory   *domain.Inventory
		quantity    int
		expectedErr error
	}{
		{
			name: "success",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  100,
			},
			quantity: 10,
		},
		{
			name: "insufficient stock",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  5,
			},
			quantity:    10,
			expectedErr: domain.ErrInsufficientStock,
		},
		{
			name:        "invalid quantity",
			quantity:    0,
			expectedErr: domain.ErrInvalidQuantity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inventoryRepo := &MockInventoryRepository{
				inventory: test.inventory,
			}

			orderRepo := &MockOrderRepository{}
			orderItemRepo := &MockOrderItemRepository{}

			service := NewOrderService(
				orderRepo,
				orderItemRepo,
				inventoryRepo,
			)

			err := service.CreateOrder(1, test.quantity, nil)

			if test.expectedErr == nil && !orderRepo.created {
				t.Fatal("expected order to be created")
			}

			if test.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if test.expectedErr != nil && !errors.Is(err, test.expectedErr) {
				t.Fatalf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func TestCreateOrderWithComment(t *testing.T) {
	comment := "urgent"

	inventoryRepo := &MockInventoryRepository{
		inventory: &domain.Inventory{
			ProductID: 1,
			Quantity:  100,
		},
	}

	orderRepo := &MockOrderRepository{}
	orderItemRepo := &MockOrderItemRepository{}

	service := NewOrderService(
		orderRepo,
		orderItemRepo,
		inventoryRepo,
	)

	err := service.CreateOrder(1, 10, &comment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if orderRepo.savedOrder == nil {
		t.Fatal("order was not saved")
	}

	if orderRepo.savedOrder.Comment == nil {
		t.Fatal("comment was not saved")
	}

	if *orderRepo.savedOrder.Comment != "urgent" {
		t.Fatalf("expected urgent, got %s", *orderRepo.savedOrder.Comment)
	}
}
