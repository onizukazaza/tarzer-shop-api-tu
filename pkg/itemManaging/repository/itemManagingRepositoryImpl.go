package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_itemManagingExceprtions "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/exception"
	_itemManagingModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/model"

)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) { 
	item := new(entities.Item)
	if err := r.db.Create(itemEntity).Scan(item).Error; err!= nil {
		r.logger.Errorf("Creating item failed: %s", err.Error()) 
		return nil, &_itemManagingExceprtions.ItemCreating{}
	} 

	return item, nil
}

func ( r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Model(&entities.Item{}).Where(
	"id =?", itemID,
	).Updates(
		itemEditingReq ,
		).Error; err!= nil {
r.logger.Errorf("Editing item failed: %s", err.Error())
return 0 ,&_itemManagingExceprtions.ItemEditing{}
		}
	return itemID, nil
}