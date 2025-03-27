package service

import (
	"github.com/google/uuid"

	"backend/internal/models"
	"backend/internal/repository"
)

type DataService struct {
	dataRepo *repository.DataRepository
}

func NewDataService(repo *repository.DataRepository) *DataService {
	return &DataService{dataRepo: repo}
}

func (s *DataService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.dataRepo.GetUser(id.String())
}

// internal/service/data.go

func (s *DataService) GetGameStatsByCode(code string) ([]models.GameUser, error) {
	return s.dataRepo.GetGameUsersByCode(code)
}
