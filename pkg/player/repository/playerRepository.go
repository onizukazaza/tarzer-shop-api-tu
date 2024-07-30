package repository

import "github.com/onizukazaza/tarzer-shop-api-tu/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player , error )
	FindByID(playerID string) (*entities.Player, error)
}