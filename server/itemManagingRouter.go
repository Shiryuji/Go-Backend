package server

import (
	_itemManagingController "github.com/Shiryuji/ryuji-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/Shiryuji/ryuji-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/Shiryuji/ryuji-shop-api/pkg/itemManaging/service"
	_itemShopRepository "github.com/Shiryuji/ryuji-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemShopRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
}