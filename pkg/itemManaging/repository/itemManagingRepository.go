package repository

import (
    "github.com/onizukazaza/tarzer-shop-api-tu/entities"
)
type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}