package controller

import (
	"net/http"
	"tutorial/pkg/custom"
	_inventoryService "tutorial/pkg/inventory/service"
	"tutorial/pkg/validation"

	"github.com/labstack/echo/v4"
)

type InventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryController(inventoryService _inventoryService.InventoryService, logger echo.Logger) InventoryController {
	return &InventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *InventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)

	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)

}
