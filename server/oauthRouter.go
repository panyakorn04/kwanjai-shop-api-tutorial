package server

import (
	_adminRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/admin/repository"
	_oauth2Controller "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/controller"
	_oauth2Service "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/service"
	_playerRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/player/repository"
)

func (s *echoServer) initOAuth2Router() {

	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _playerRepository.NewPlayerRepository(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		s.app.Logger,
	)

	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/player/login/callback", oauth2Controller.PlayerCallback)
	router.GET("/admin/login/callback", oauth2Controller.AdminCallback)
	router.DELETE("/logout", oauth2Controller.Logout)

}
