package service

import (
	_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
)
type ItemShopService interface{
	Listing(itemFilter *_itemShopModel.ItemFilter) ( *_itemShopModel.ItemResult, error) 
}

