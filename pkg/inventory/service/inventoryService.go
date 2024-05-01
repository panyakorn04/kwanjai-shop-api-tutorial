package service

import _inventoryModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/model"

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
