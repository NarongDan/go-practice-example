package service

import (
	"tutorial/entities"
	_investoryModel "tutorial/pkg/inventory/model"
	_inventoryRepository "tutorial/pkg/inventory/repository"
	_itemShopRepository "tutorial/pkg/itemShop/repository"
)

type InventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(inventoryRepository _inventoryRepository.InventoryRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository) InventoryService {
	return &InventoryServiceImpl{inventoryRepository: inventoryRepository,
		itemShopRepository: itemShopRepository}
}

func (s *InventoryServiceImpl) Listing(playerID string) ([]*_investoryModel.Inventory, error) {

	inventoryEntities, err := s.inventoryRepository.Listing(playerID)

	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventoryEntities)

	return s.buildInventoryListingResult(uniqueItemWithQuantityCounterList), nil

}

func (s *InventoryServiceImpl) getUniqueItemWithQuantityCounterList(
	inventoryEntities []*entities.Inventory,
) []_investoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_investoryModel.ItemQuantityCounting, 0)
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _investoryModel.ItemQuantityCounting{ItemID: itemID, Quantity: quantity})
	}

	return itemQuantityCounterList

}
func (s *InventoryServiceImpl) buildInventoryListingResult(
	uniqueItemWithQuantityCounterList []_investoryModel.ItemQuantityCounting,
) []*_investoryModel.Inventory {

	uniqueItemIDList := s.getItemID(uniqueItemWithQuantityCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)

	if err != nil {
		return make([]*_investoryModel.Inventory, 0)
	}

	results := make([]*_investoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapIWthQuantity(uniqueItemWithQuantityCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_investoryModel.Inventory{Item: itemEntity.ToItemModel(),

			Quanity: itemMapWithQuantity[itemEntity.ID]},
		)
	}

	return results

}

func (s *InventoryServiceImpl) getItemID(uniqueItemWithQuantityCounterList []_investoryModel.ItemQuantityCounting) []uint64 {

	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *InventoryServiceImpl) getItemMapIWthQuantity(
	uniqueItemWithQuantityCounterList []_investoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
