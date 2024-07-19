package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_itemManagingExceprtions "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/exception"

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
// func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
// 	if err := r.db.Create(itemEntity).Error; err!= nil {
// 		return nil , err
// 	}
// 	return itemEntity, nil
// }