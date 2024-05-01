package repository

import (
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int, tx *gorm.DB) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
