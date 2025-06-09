package server

import (
	_inventoryController "tutorial/pkg/inventory/controller"
	_inventoryRepository "tutorial/pkg/inventory/repository"
	_inventoryService "tutorial/pkg/inventory/service"
	_itemShopRepository "tutorial/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	itemShopRepository := _itemShopRepository.NewItemShopRepository(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(inventoryRepository, itemShopRepository)

	inventoryController := _inventoryController.NewInventoryController(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing)

}
