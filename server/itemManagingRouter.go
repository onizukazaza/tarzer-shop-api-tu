package server

import ( 
	_itemManagingController "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/repository"
    _itemManagingService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/service"
	_itemShopRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/repository"
)
func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(
		itemManagingRepository,
		itemShopRepository,
	)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
	router.DELETE("/:itemID", itemManagingController.Archiving)
}