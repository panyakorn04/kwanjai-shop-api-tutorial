package server

import (
	_playerCoinController "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/repository"
	_playerCoinService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddlewares) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorize)
	router.GET("", playerCoinController.Showing, m.PlayerAuthorize)

}
