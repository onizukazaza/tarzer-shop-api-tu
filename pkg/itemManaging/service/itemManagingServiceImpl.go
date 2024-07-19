package service

import (
_itemManagingRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/repository"
)

type  ItemManagingServiceImpl  struct {
   itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	)ItemManagingService{
   return &ItemManagingServiceImpl{itemManagingRepository}
}
