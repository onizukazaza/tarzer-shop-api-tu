package service 

import (
	_playerModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/model"
	_adminModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyReallyPlayer(playerID string) bool
	IsThisGuyReallyAdmin(adminID string) bool
}