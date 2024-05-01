package server

import (
	_itemShopRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/repository"

	_inventoryController "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/controller"
	_inventoryRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/repository"
	_inventoryService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/service"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddlewares) {
	router := s.app.Group("/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(inventoryRepository, itemShopRepository)

	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorize)
}
