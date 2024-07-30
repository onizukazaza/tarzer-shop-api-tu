package service

import (
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_adminModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/model"
	_adminRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/repository"
	_playerModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/model"
	_playerRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
} //build player and admin input databases

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}
func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	if !s.IsThisGuyReallyPlayer(playerCreatingReq.ID) {
	playerEnitity := &entities.Player{
		ID:     playerCreatingReq.ID,
		Name:   playerCreatingReq.Name,
		Email:  playerCreatingReq.Email,
		Avatar: playerCreatingReq.Avatar,
	}
	if _, err := s.playerRepository.Creating(playerEnitity); err != nil {
		return err
	}
}
	return nil
}
func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {
	if !s.IsThisGuyReallyAdmin(adminCreatingReq.ID) {
		adminEnitity := &entities.Admin{
			ID:     adminCreatingReq.ID,
			Name:   adminCreatingReq.Name,
			Email:  adminCreatingReq.Email,
			Avatar: adminCreatingReq.Avatar,
		}
		if _, err := s.adminRepository.Creating(adminEnitity); err != nil {
			return err
		}
	}
		return nil
	}
func (s *googleOAuth2Service) IsThisGuyReallyPlayer(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err!= nil {
		return false
	}
	return player!= nil
}
 

func (s *googleOAuth2Service) IsThisGuyReallyAdmin(adminID string) bool {
	admin, err := s.playerRepository.FindByID(adminID)
	if err!= nil {
		return false
	}
	return admin!= nil
}
 

