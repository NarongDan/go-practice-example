package service

import (
	"tutorial/entities"
	_adminModel "tutorial/pkg/admin/model"
	_adminRepository "tutorial/pkg/admin/repository"
	_playerModel "tutorial/pkg/player/model"
	_playerRepository "tutorial/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository) OAuth2Service {

	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {

	if !s.IsThisGuyIsReallyPlayer(playerCreatingReq.ID) {

		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Email:  playerCreatingReq.Email,
			Name:   playerCreatingReq.Name,
			Avatar: playerCreatingReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {

	if !s.IsThisGuyIsReallyAdmin(adminCreatingReq.ID) {

		adminEntity := &entities.Admin{
			ID:     adminCreatingReq.ID,
			Email:  adminCreatingReq.Email,
			Name:   adminCreatingReq.Name,
			Avatar: adminCreatingReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {

	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {

	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}

	return admin != nil
}
