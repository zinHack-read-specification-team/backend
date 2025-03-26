package repository

import (
	"backend/internal/models"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameLinkRepository struct {
	db *gorm.DB
}

func NewGameLinkRepository(db *gorm.DB) *GameLinkRepository {
	return &GameLinkRepository{db: db}
}

// CreateGameLink создает новую ссылку
func (r *GameLinkRepository) CreateGameLink(link *models.GameLink) error {
	for i := 0; i < 5; i++ {
		link.Code = generateCode(6)

		// сохраняем ошибку в переменную
		err := r.db.Create(link).Error
		if err == nil {
			return nil
		}
		// если ошибка НЕ про дубликат — сразу возвращаем
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return err
		}
	}

	return errors.New("failed to generate unique code after several attempts")
}

func generateCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seed.Intn(len(charset))]
	}
	return string(code)
}

func (r *GameLinkRepository) GetLinksByUserID(userID uuid.UUID) ([]models.GameLink, error) {
	var links []models.GameLink
	if err := r.db.Where("user_id = ?", userID).Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}
