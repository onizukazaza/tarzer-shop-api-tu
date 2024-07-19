package server

import (
	_itemShopRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/repository"
	_itemShopService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/service"
	_itemShopController "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/controller"
)
func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("" , itemShopController.Listing )
}