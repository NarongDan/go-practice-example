package server

import (
	_inventoryRepository "tutorial/pkg/inventory/repository"
	_itemShopController "tutorial/pkg/itemShop/controller"
	_itemShopRepository "tutorial/pkg/itemShop/repository"
	_itemShopService "tutorial/pkg/itemShop/service"
	_playerCoinRepository "tutorial/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")
	itemShopRepository := _itemShopRepository.NewItemShopRepository(s.db, s.app.Logger)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopService(itemShopRepository, playerCoinRepository, inventoryRepository, s.app.Logger)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("", itemShopController.Buying, m.PlayerAuthorizing)
	router.POST("", itemShopController.Selling, m.PlayerAuthorizing)
}
