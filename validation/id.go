package validation

import (
	"github.com/labstack/echo/v4"
	_admin "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/admin/exception"
	_player "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/player/exception"
)

func AdminIDGetting(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_admin.AdminNotFound{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}

func PlayerIDGetting(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_player.PlayerNotFound{PlayerID: "Unknown"}
	} else {
		return playerID, nil
	}
}
