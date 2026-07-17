package service

import (
	"errors"
	"testing"

	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

var _ repository.ProductRepository = (*MockProductRepository)(nil)

type MockProductRepository struct {
	product *domain.Product
	err     error
}

func (m *MockProductRepository) Create(
	product *domain.Product,
) error {
	return nil
}

func (m *MockProductRepository) List() ([]domain.Product, error) {
	return nil, nil
}

func (m *MockInventoryRepository) Create(
	productID int64,
	quantity int,
) error {
	return nil
}

func (m *MockInventoryRepository) DecreaseStock(
	productID int64,
	quantity int,
) error {
	return nil
}

func (m *MockProductRepository) GetByID(id int64) (*domain.Product, error) {
	return m.product, m.err
}

func TestGetProduct(t *testing.T) {
	expected := &domain.Product{
		ID:   1,
		Name: "Mug",
	}

	repo := &MockProductRepository{
		product: expected,
	}

	service := NewProductService(repo, &MockInventoryRepository{})

	actual, err := service.GetProduct(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual.Name != expected.Name {
		t.Fatalf("expected %s, got %s", expected.Name, actual.Name)
	}
}

func TestGetProductTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		product     *domain.Product
		err         error
		expectedErr error
	}{
		{
			name: "product exists",
			product: &domain.Product{
				ID:   1,
				Name: "Mug",
			},
			expectedErr: nil,
		},
		{
			name:        "product not found",
			err:         domain.ErrProductNotFound,
			expectedErr: domain.ErrProductNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &MockProductRepository{
				product: test.product,
				err:     test.err,
			}

			service := NewProductService(repo, &MockInventoryRepository{})

			_, err := service.GetProduct(1)

			if test.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if test.expectedErr == nil {
				return
			}

			if !errors.Is(err, test.expectedErr) {
				t.Fatalf("expected %v, got %v", test.expectedErr, err)
			}

		})
	}
}
