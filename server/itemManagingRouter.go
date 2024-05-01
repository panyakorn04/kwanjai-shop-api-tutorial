package server

import (
	_itemManagingController "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/repository"
	_itemManagingService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/service"
	_itemShopRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddlewares) {
	router := s.app.Group("/v1/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemRepository)
	itemManaging := _itemManagingController.NewItemManagingController(itemManagingService)

	router.POST("", itemManaging.Creating, m.AdminAuthorize)
	router.PATCH("/:itemID", itemManaging.Editing, m.AdminAuthorize)
	router.DELETE("/:itemID", itemManaging.Archiving, m.AdminAuthorize)
}
