package service

import (
	"tutorial/entities"
	_playerCoinException "tutorial/pkg/playerCoin/exception"
	_playerCoinModel "tutorial/pkg/playerCoin/model"
	_playerCoinRepository "tutorial/pkg/playerCoin/repository"
)

type PlayerCoinServiceImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinService(playerCoinRepository _playerCoinRepository.PlayerCoinRepository) PlayerCoinService {
	return &PlayerCoinServiceImpl{playerCoinRepository: playerCoinRepository}
}

func (s *PlayerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerEntityResult, err := s.playerCoinRepository.CoinAdding(nil, playerCoinEntity)

	if err != nil {
		return nil, &_playerCoinException.CoinAdding{}
	}

	playerEntityResult.PlayerID = coinAddingReq.PlayerID

	return playerEntityResult.ToPlayerCoinModel(), nil
}

func (s *PlayerCoinServiceImpl) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)

	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}

	return playerCoinShowing

}
