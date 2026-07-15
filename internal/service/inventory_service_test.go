package service

import (
	"errors"
	"testing"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.InventoryRepository = (*MockInventoryRepository)(nil)

type MockInventoryRepository struct {
	inventory *domain.Inventory
	err       error
}

func (m *MockInventoryRepository) GetByProductID(productID int64) (*domain.Inventory, error) {
	return m.inventory, m.err
}

func (m *MockInventoryRepository) UpdateQuantity(productID int64, quantity int) error {
	return nil
}
func TestGetStock(t *testing.T) {
	repo := &MockInventoryRepository{
		inventory: &domain.Inventory{
			ProductID: 1,
			Quantity:  100,
		},
	}

	service := NewInventoryService(repo)

	stock, err := service.GetStock(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stock != 100 {
		t.Fatalf("expected 100, got %d", stock)
	}
}

func TestGetStockTableDriven(t *testing.T) {
	tests := []struct {
		name          string
		inventory     *domain.Inventory
		err           error
		expectedStock int
		expectedErr   error
	}{
		{
			name: "stock exists",
			inventory: &domain.Inventory{
				ProductID: 1,
				Quantity:  100,
			},
			expectedStock: 100,
		},
		{
			name:        "repository error",
			err:         domain.ErrProductNotFound,
			expectedErr: domain.ErrProductNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &MockInventoryRepository{
				inventory: test.inventory,
				err:       test.err,
			}

			service := NewInventoryService(repo)

			stock, err := service.GetStock(1)

			if test.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if test.expectedErr != nil {
				if !errors.Is(err, test.expectedErr) {
					t.Fatalf("expected %v, got %v", test.expectedErr, err)
				}
				return
			}

			if stock != test.expectedStock {
				t.Fatalf("expected %d, got %d", test.expectedStock, stock)
			}
		})
	}
}
