package controller

import "github.com/labstack/echo/v4"

type OAuth2Controller interface {
	PlayerLogin(pctx echo.Context) error
	AdminLogin(pctx echo.Context) error
	PlayerCallback(pctx echo.Context) error
	AdminCallback(pctx echo.Context) error
	Logout(pctx echo.Context) error

	PlayerAuthorize(pctx echo.Context, next echo.HandlerFunc) error
	AdminAuthorize(pctx echo.Context, next echo.HandlerFunc) error
}
