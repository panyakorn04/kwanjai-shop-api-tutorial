package service

import (
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
}
