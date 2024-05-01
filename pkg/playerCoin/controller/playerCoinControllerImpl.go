package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/custom"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
	_playerCoinService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/service"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/validation"
)

type playerCoinController struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinController{
		playerCoinService: playerCoinService,
	}
}

func (c *playerCoinController) CoinAdding(pctx echo.Context) error {
	plyerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoReq := custom.NewCustomEchoRequest(pctx)
	if err := customEchoReq.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq.PlayerID = plyerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinController) Showing(pctx echo.Context) error {
	plyerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoinShowing := c.playerCoinService.Showing(plyerID)

	return pctx.JSON(http.StatusOK, playerCoinShowing)

}
