package service

import (
	_playerCoinModel "tutorial/pkg/playerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	Showing(playerID string) *_playerCoinModel.PlayerCoinShowing
}
