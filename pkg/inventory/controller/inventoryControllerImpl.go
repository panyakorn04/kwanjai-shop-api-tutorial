package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/custom"
	_inventoryService "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/service"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/validation"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryControllerImpl(inventoryService _inventoryService.InventoryService, logger echo.Logger) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		c.logger.Errorf("error getting player id: %s", err)
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	inventories, err := c.inventoryService.Listing(playerID)
	if err != nil {
		c.logger.Errorf("error listing inventory: %s", err)
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventories)
}
