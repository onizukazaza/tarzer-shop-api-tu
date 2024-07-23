package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_itemShopException "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/exception"
	_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"

)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0) //type  pointer first len 0

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
	// 1 2 3 4 5 6 7 8 9 10
	//0        |5        |
	//(page - 1) *size/limit
	//(1 - 1) *5 = 0

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(int(limit)).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Errorf("Failed to list item: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}
	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
//    count := new(int64)
   var count int64

	if err :=query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting item failed: %s", err.Error())
		return -1, &_itemShopException.ItemCounting{}
	}
	return count, nil
}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item , error) {
 item := new(entities.Item)

 if err := r.db.Connect().First(item, itemID).Error; err!= nil {
    r.logger.Errorf("Failed to find item by ID: %s", err.Error())
    return nil, &_itemShopException.ItemNotFound{}
 }
 return item, nil
}