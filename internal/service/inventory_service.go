package service

import "github.com/kpiasecki/wms/internal/repository"

type InventoryService struct {
	inventoryRepository repository.InventoryRepository
}

func NewInventoryService(
	inventoryRepository repository.InventoryRepository,
) *InventoryService {
	return &InventoryService{
		inventoryRepository: inventoryRepository,
	}
}

func (s *InventoryService) GetStock(productID int64) (int, error) {
	inventory, err := s.inventoryRepository.GetByProductID(productID)
	if err != nil {
		return 0, err
	}

	return inventory.Quantity, nil
}
