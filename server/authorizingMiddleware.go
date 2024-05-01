package server

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/config"
	_oauth2Controller "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/controller"
)

type authorizingMiddlewares struct {
	OAuth2Controller _oauth2Controller.OAuth2Controller
	oauth2Conf       *config.OAuth2
	logger           echo.Logger
}

func (m *authorizingMiddlewares) PlayerAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.OAuth2Controller.PlayerAuthorize(pctx, next)
	}
}

func (m *authorizingMiddlewares) AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.OAuth2Controller.AdminAuthorize(pctx, next)
	}
}
