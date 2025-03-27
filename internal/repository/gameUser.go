// repository/game_user.go
package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type GameUserRepository struct {
	db *gorm.DB
}

func NewGameUserRepository(db *gorm.DB) *GameUserRepository {
	return &GameUserRepository{db: db}
}

func (r *GameUserRepository) CreateGameUser(user *models.GameUser) error {
	return r.db.Create(user).Error
}

func (r *GameUserRepository) UpdateScore(id string, stars, score int) error {
	return r.db.Model(&models.GameUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"stars": stars,
			"score": score,
		}).Error
}

func (r *GameLinkRepository) GetGameLinkByCode(code string) (*models.GameLink, error) {
	var link models.GameLink
	if err := r.db.Where("code = ?", code).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}
