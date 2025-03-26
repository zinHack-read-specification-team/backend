package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/google/uuid"
)

type GameLinkService struct {
	repo *repository.GameLinkRepository
}

func NewGameLinkService(r *repository.GameLinkRepository) *GameLinkService {
	return &GameLinkService{repo: r}
}

func (s *GameLinkService) CreateLink(userID uuid.UUID, gameName, schoolNum, class, coment string) (*models.GameLink, error) {
	link := &models.GameLink{
		UserID:    userID,
		GameName:  gameName,
		SchoolNum: schoolNum,
		Class:     class,
		Comment:   coment,
	}
	err := s.repo.CreateGameLink(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
func (s *GameLinkService) GetLinksByUserID(userID uuid.UUID) ([]models.GameLink, error) {
	return s.repo.GetLinksByUserID(userID)
}
