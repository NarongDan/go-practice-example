package repository

import (
	"tutorial/databases"
	"tutorial/entities"

	_inventoryException "tutorial/pkg/inventory/exception"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type InventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &InventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *InventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{PlayerID: playerID, ItemID: itemID})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("Filling inventory failed: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{ItemID: itemID, PlayerID: playerID}
	}

	return inventoryEntities, nil
}

func (r *InventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)

	if err != nil {
		return err
	}

	// tx := r.db.Connect().Begin()

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := conn.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("Removing player item in inventory failed: %s", err.Error())
			return &_inventoryException.PlayerItemRemoving{ItemID: itemID}

		}

	}

	// if err := tx.Commit().Error; err != nil {
	// 	tx.Rollback()
	// 	r.logger.Errorf("Removing player item in inventory failed: %s", err.Error())
	// 	return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
	// }

	return nil
}

func (r *InventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id  = ? and item_id = ? and is_deleted = ?", playerID, itemID, false).Count(&count).Error; err != nil {
		r.logger.Errorf("Counting player item in inventory failed: %s", err.Error())
		return -1
	}

	return count
}

func (r *InventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id  = ? and is_deleted = ?", playerID, false).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("Listing player item in inventory failed: %s", err.Error())
		return nil, &_inventoryException.PlayerItemsFinding{}
	}

	return inventoryEntities, nil
}

func (r *InventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id  = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("Find player item in inventory by id failed: %s", err.Error())
		return nil, &_inventoryException.PlayerItemRemoving{ItemID: itemID}
	}

	return inventoryEntities, nil

}
