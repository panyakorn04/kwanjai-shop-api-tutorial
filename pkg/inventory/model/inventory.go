package model

import (
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
)

type (
	Inventory struct {
		Item     *_itemShopModel.Item `json:"item"`
		Quantity uint                 `json:"quality"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
