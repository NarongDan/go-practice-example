package server

import (
	_playerCoinController "tutorial/pkg/playerCoin/controller"
	_playerCoinRepository "tutorial/pkg/playerCoin/repository"
	_playerCoinService "tutorial/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {

	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	playerCoinService := _playerCoinService.NewPlayerCoinService(playerCoinRepository)

	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
	router.GET("", playerCoinController.Showing, m.PlayerAuthorizing)

}
