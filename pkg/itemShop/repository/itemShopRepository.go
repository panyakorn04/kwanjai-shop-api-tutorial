package repository

import (
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	TransactionBegin() *gorm.DB
	TransactionCommit(tx *gorm.DB) error
	TransactionRollback(tx *gorm.DB) error
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDs []uint64) ([]entities.Item, error)
	PurchaseHistoryRecording(
		purchasingEntity *entities.PurchaseHistory,
		tx *gorm.DB,
	) (*entities.PurchaseHistory, error)
}
