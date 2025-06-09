package repository

import (
	"tutorial/databases"
	"tutorial/entities"

	_playerCoinException "tutorial/pkg/playerCoin/exception"
	_playerCoinModel "tutorial/pkg/playerCoin/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PlayerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &PlayerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *PlayerCoinRepositoryImpl) CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {

	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	playerCoin := new(entities.PlayerCoin)

	if err := conn.Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Errorf("Adding player coin failed: %s", err.Error())
		return nil, &_playerCoinException.CoinAdding{}
	}

	return playerCoin, nil

}

func (r *PlayerCoinRepositoryImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {

	playerCoinShowing := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(&entities.PlayerCoin{}).Where(
		"player_id =?", playerID,
	).Select("player_id, sum(amount) as coin").Group("player_id").Scan(playerCoinShowing).Error; err != nil {
		r.logger.Errorf("Showing player coin failed: %s", err.Error())
		return nil, &_playerCoinException.PLayerCoinShowing{}
	}

	return playerCoinShowing, nil

}
