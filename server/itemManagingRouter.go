package server

import ( 
	_itemManagingController "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/repository"
    _itemManagingService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/service"
)
func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-Managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)
	
	_ = itemManagingController
	_ = router
}