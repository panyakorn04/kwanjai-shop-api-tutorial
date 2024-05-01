package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_itemShop "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/exception"
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemShopRepositoryImpl) TransactionBegin() *gorm.DB {
	tx := r.db.Connect().Begin()
	return tx.Begin()
}

func (r *itemShopRepositoryImpl) TransactionCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *itemShopRepositoryImpl) TransactionRollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]entities.Item, error) {
	itemList := make([]entities.Item, 0)
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)
	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Order("id desc").Error; err != nil {
		r.logger.Error("Error while fetching items", err)
		return nil, &_itemShop.ItemListing{}
	}

	return itemList, nil

}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)
	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Error while counting items", err)
		return -1, &_itemShop.ItemCounting{}
	}

	return count, nil

}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Error("Error while fetching item by id", err)
		return nil, &_itemShop.ItemNotFound{}
	}
	return item, nil
}

func (r *itemShopRepositoryImpl) FindByIDList(itemIDs []uint64) ([]entities.Item, error) {
	itemList := make([]entities.Item, 0)

	if err := r.db.Connect().Model(&entities.Item{}).Where("id IN ?", itemIDs).Find(&itemList).Error; err != nil {
		r.logger.Error("Error while fetching item by id list", err)
		return nil, &_itemShop.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) PurchaseHistoryRecording(
	purchasingEntity *entities.PurchaseHistory,
	tx *gorm.DB,
) (*entities.PurchaseHistory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	insertedPurchasing := new(entities.PurchaseHistory)

	if err := conn.Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Purchase history recording failed: %s", err.Error())
		return nil, &_itemShop.HistoryOfPurchaseRecording{}
	}

	return insertedPurchasing, nil
}
