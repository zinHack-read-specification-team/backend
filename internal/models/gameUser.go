package models

import (
	"time"

	"github.com/google/uuid"
)

// @Description Структура, описывающая пользователя
type GameUser struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID    uuid.UUID `json:"game_id" gorm:"type:uuid;not null"`
	GameName  string    `json:"game_name"`
	GameCode  string    `json:"game_code"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Stars     int       `json:"stars"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
