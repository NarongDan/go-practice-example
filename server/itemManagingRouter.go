package server

import (
	_itemManaingController "tutorial/pkg/itemManaging/controller"
	_itemManagingRepository "tutorial/pkg/itemManaging/repository"
	_itemManagingService "tutorial/pkg/itemManaging/service"
	_itemShopRepository "tutorial/pkg/itemShop/repository"
)

func (s *echoServer) initItemManaingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepository(s.db, s.app.Logger)

	itemShopRepository := _itemShopRepository.NewItemShopRepository(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemShopRepository)

	itemManagingController := _itemManaingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating, m.AdminAuthorizing)
	router.PATCH("/:itemID", itemManagingController.Editing, m.AdminAuthorizing)
	router.DELETE("/:itemID", itemManagingController.Archiving, m.AdminAuthorizing)
	// router.GET("", itemShopController.Listing)
}
