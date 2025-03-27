// service/game_user.go
package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/google/uuid"
)

type GameUserService struct {
	repo         *repository.GameUserRepository
	gameLinkRepo *repository.GameLinkRepository
}

func NewGameUserService(r *repository.GameUserRepository, g *repository.GameLinkRepository) *GameUserService {
	return &GameUserService{repo: r, gameLinkRepo: g}
}

func (s *GameUserService) RegisterPlayer(code, fullName string) (*models.GameUser, error) {
	game, err := s.gameLinkRepo.GetGameLinkByCode(code)
	if err != nil {
		return nil, err
	}

	user := &models.GameUser{
		GameID:   game.ID,
		GameName: game.GameName,
		GameCode: game.Code,
		FullName: fullName,
		Stars:    0,
		Score:    0,
	}
	if err := s.repo.CreateGameUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *GameUserService) FinishGame(id uuid.UUID, stars, score int) error {
	return s.repo.UpdateScore(id.String(), stars, score)
}
