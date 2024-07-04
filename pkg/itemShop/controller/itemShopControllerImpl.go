package controller

import(
	_itemShopService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(	itemShopService _itemShopService.ItemShopService ,
) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}


