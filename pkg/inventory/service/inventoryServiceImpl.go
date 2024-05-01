package service

import (
	entities "github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_inventoryModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/model"
	_inventoryRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/repository"
	_itemShopRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(inventoryRepository _inventoryRepository.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		itemShopRepository:  itemShopRepository,
	}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventories, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventories)
	return s.buildInventoryListingResult(uniqueItemWithQuantityCounterList), nil
}

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(inventoryEntities []*entities.Inventory) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}
	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) buildInventoryListingResult(itemQuantityCounterList []_inventoryModel.ItemQuantityCounting) []*_inventoryModel.Inventory {
	itemIDList := s.getItemID(itemQuantityCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(itemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	result := make([]*_inventoryModel.Inventory, 0)

	itemMapWithQuantity := s.getItemMapWithQuantity(itemQuantityCounterList)

	for _, itemEntity := range itemEntities {
		result = append(result, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}

	return result
}

func (s *inventoryServiceImpl) getItemID(
	itemQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range itemQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(
	itemQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range itemQuantityCounterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
