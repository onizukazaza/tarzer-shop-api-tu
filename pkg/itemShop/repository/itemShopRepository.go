package repository

import (
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_itemShopModel"github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
) 
type ItemShopRepository interface{

Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
}