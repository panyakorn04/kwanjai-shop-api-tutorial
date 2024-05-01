package server

import (
	_inventoryRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/repository"
	_itemShopController "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/controller"
	_itemShopRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/repository"
	_itemShopService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/service"
	_playerCoinRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddlewares) {
	router := s.app.Group("/v1/item-shop")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository, playerCoinRepository, inventoryRepository, s.app.Logger)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorize)
	router.POST("/selling", itemShopController.Selling)
}
