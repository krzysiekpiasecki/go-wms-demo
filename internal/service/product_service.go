package service

import (
	"github.com/kpiasecki/wms/internal/domain"
	"github.com/kpiasecki/wms/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(
	productRepository repository.ProductRepository,
) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	return s.productRepository.GetByID(id)
}
