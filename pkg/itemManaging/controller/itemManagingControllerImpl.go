package controller

import (
	_itemManagingService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/service" 
)
type ItemManagingControllerImpl struct {
	ItemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(
	itemManagingService _itemManagingService.ItemManagingService,
    ) ItemManagingController {
    return &ItemManagingControllerImpl{itemManagingService}
	}
