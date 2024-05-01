package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/custom"
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
	_itemShopService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/service"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/validation"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{
		itemShopService: itemShopService,
	}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	sellingReq := new(_itemShopModel.SellingReq)

	validatingContext := custom.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(sellingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID

	result, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	buyingReq := new(_itemShopModel.BuyingReq)

	validatingContext := custom.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID

	result, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}
