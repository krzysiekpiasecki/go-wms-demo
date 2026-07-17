package service

import (
	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

type ProductService struct {
	productRepository   repository.ProductRepository
	inventoryRepository repository.InventoryRepository
}

func NewProductService(
	productRepository repository.ProductRepository,
	inventoryRepository repository.InventoryRepository,
) *ProductService {
	return &ProductService{
		productRepository:   productRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.productRepository.List()
}

func (s *ProductService) CreateProduct(
	name string,
	quantity int,
) (*domain.Product, error) {

	product := &domain.Product{
		Name: name,
	}

	err := s.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	err = s.inventoryRepository.Create(
		product.ID,
		quantity,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
