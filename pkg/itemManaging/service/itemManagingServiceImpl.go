package service

import (
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_itemManagingModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/model"
	_itemManagingRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/repository"
	_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
	_itemShopRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/repository"
)

type ItemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &ItemManagingServiceImpl{
		itemManagingRepository,
		itemShopRepository,
	}
}

// func NewItemManagingServiceImpl(
//
//	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
//	itemShopRepository _itemShopRepository.ItemShopRepository,
//
//	) ItemManagingService {
//		return &ItemManagingServiceImpl{
//			itemManagingRepository: itemManagingRepository,
//			itemShopRepository:     itemShopRepository,
//		}
//	}
func (s *ItemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}
	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}
	return itemEntityResult.ToItemModel(), nil
}

func (s *ItemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	_, err := s.itemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepository.FindByID(itemID)
	if err != nil {
		return nil, err
	}
	return itemEntityResult.ToItemModel(), nil
}

func (s *ItemManagingServiceImpl) Archiving(itemID uint64) error {
	return s.itemManagingRepository.Archiving(itemID)
}