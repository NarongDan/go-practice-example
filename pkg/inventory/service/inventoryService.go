package service

import (
	_investoryModel "tutorial/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_investoryModel.Inventory, error)
}
