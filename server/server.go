package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/config"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"

	_adminRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/admin/repository"
	_oauth2Controller "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/controller"
	_oauth2Service "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/service"
	_playerRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/player/repository"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *echoServer) Start() {
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	getTimeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(getTimeOutMiddleware)

	s.app.GET("/v1/health", s.healthCheck)

	authorizingMiddlewares := s.getAuthorizationMiddleware()

	s.initItemShopRouter(authorizingMiddlewares)
	s.initItemManagingRouter(authorizingMiddlewares)
	s.initOAuth2Router()
	s.initPlayerCoinRouter(authorizingMiddlewares)
	s.initInventoryRouter(authorizingMiddlewares)
	// Graceful shutdown
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)
	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("shutting down the server: %v", err)
	}
}

func (s *echoServer) gracefullyShutdown(quitCh <-chan os.Signal) {
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("shutting down the server")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("shutting down the server: %v", err)
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "I'm alive")
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout:      timeout * time.Second,
	})
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleware(limit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(limit)
}

func (s *echoServer) getAuthorizationMiddleware() *authorizingMiddlewares {
	playerRepository := _playerRepository.NewPlayerRepository(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Server := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Server,
		s.conf.OAuth2,
		s.app.Logger,
	)

	return &authorizingMiddlewares{
		OAuth2Controller: oauth2Controller,
		oauth2Conf:       s.conf.OAuth2,
		logger:           s.app.Logger,
	}
}
