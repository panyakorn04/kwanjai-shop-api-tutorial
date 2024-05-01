package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_itemManagingException "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/exception"
	_itemManagingModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/model"
)

type ItemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) *ItemManagingRepositoryImpl {
	return &ItemManagingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *ItemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Error creating item: %v", err)
		return nil, &_itemManagingException.ItemCreating{}

	}
	return item, nil
}

func (r *ItemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		itemEditingReq,
	).Error; err != nil {
		r.logger.Error("Editing item failed:", err)
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *ItemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Error("Archiving item failed:", err)
		return &_itemManagingException.ItemArchiving{}
	}
	return nil
}
