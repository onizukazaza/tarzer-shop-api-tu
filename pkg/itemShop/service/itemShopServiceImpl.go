package service

	 import (
		_itemShopRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/repository"
		_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
		"github.com/onizukazaza/tarzer-shop-api-tu/entities"
		
		
	)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository //_itemShopRepository นำเข้า
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository ,
	) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter) //inject itemshop
	if err != nil {
		return nil, err
	}

itemCounting, err := s.itemShopRepository.Counting(itemFilter)
if err != nil {
	return nil, err
}

size := itemFilter.Size
page := itemFilter.Page
totalPage := s.totalPageCalculation(itemCounting,size)

result := s.toItemResultRespone(itemList, page, totalPage)

	return result,nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size // 11 / 5 = 2

	if totalItems%size != 0 {
		totalPage++
	}

return totalPage
}

	func (s *itemShopServiceImpl) toItemResultRespone(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
		itemModelList := make([]*_itemShopModel.Item, 0)
		for _, item := range itemEntityList {
			itemModelList = append(itemModelList, item.ToItemModel())
		} 
		return &_itemShopModel.ItemResult{
			Items: itemModelList,
			Paginate: _itemShopModel.PaginateResult{
			Page: page,
			Totalpage: totalPage,
		},
		}
}