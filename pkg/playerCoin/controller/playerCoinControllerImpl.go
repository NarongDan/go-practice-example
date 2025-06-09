package controller

import (
	"net/http"
	"tutorial/pkg/custom"
	_playerCoinModel "tutorial/pkg/playerCoin/model"
	_playerCoinService "tutorial/pkg/playerCoin/service"
	"tutorial/pkg/validation"

	"github.com/labstack/echo/v4"
)

type PlayerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &PlayerCoinControllerImpl{playerCoinService}
}

func (c *PlayerCoinControllerImpl) CoinAdding(pctx echo.Context) error {

	// playerID, err := validation.PlayerIDGetting(pctx)

	// if err != nil {
	// 	return custom.Error(pctx, http.StatusBadRequest, err)
	// }

	playerID, err := validation.PlayerIDGetting(pctx)
	pctx.Logger().Infof("playerID is ---------%v", playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *PlayerCoinControllerImpl) Showing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	pctx.Logger().Infof("playerID is ---------%v", playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoinShowing := c.playerCoinService.Showing(playerID)

	return pctx.JSON(http.StatusOK, playerCoinShowing)
}
