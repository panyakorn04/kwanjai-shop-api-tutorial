package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/custom"
	_oauth2Exception "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/oauth2/exception"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorize(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()
	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.PlayerTokenRefresh(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsThisGuyIsReallyPlayer(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) AdminAuthorize(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()
	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.AdminTokenRefresh(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsThisGuyIsReallyAdmin(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Unauthorized{})
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) PlayerTokenRefresh(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil

}

func (c *googleOAuth2Controller) AdminTokenRefresh(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	refreshToken, err := pctx.Cookie(refreshTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}
	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
