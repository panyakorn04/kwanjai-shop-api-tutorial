package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_inventory "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/exception"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("error filling player item in inventory: %s", err.Error())
		return nil, &_inventory.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(playerID string, itemID uint64, limit int, tx *gorm.DB) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := conn.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s", err.Error())
			return &_inventory.PlayerItemRemoving{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false).Count(&count).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory: %v", err.Error())
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id = ? AND is_deleted = ?", playerID, false).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error listing inventory: %s", err.Error())
		return nil, &_inventory.InventoryFilling{PlayerID: playerID}
	}

	return inventoryEntities, nil

}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error finding player item in inventory: %s", err.Error())
		return nil, &_inventory.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventoryEntities, nil
}
