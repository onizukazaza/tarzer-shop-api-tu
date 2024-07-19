package service

import (
_itemManagingRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/repository"
_itemManagingModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/model"
_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
    "github.com/onizukazaza/tarzer-shop-api-tu/entities"
)

type  ItemManagingServiceImpl  struct {
   itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	)ItemManagingService{
   return &ItemManagingServiceImpl{itemManagingRepository}
}

func (s *ItemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name: itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
        Picture: itemCreatingReq.Picture,
        Price: itemCreatingReq.Price,
	}
	itemEntityResult , err := s.itemManagingRepository.Creating(itemEntity)
	if err!= nil {
        return nil, err
    }
	return itemEntityResult.ToItemModel(), nil
}