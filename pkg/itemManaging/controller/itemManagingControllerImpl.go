package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	custom "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/custom"
	_itemManagingModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/model"
	_itemManaging "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemManaging/service"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/validation"
)

type itemManagingImpl struct {
	itemManaging _itemManaging.ItemManagingService
}

func NewItemManagingController(itemManaging _itemManaging.ItemManagingService) ItemManagingController {
	return &itemManagingImpl{itemManaging: itemManaging}
}

func (c *itemManagingImpl) Creating(pctx echo.Context) error {
	adminID, err := validation.AdminIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.itemManaging.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingImpl) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	itemEditingReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemEditingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManaging.Editing(itemID, itemEditingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")

	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}

func (c *itemManagingImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	if err := c.itemManaging.Archiving(itemID); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}
