package models

import (
	"time"

	"github.com/google/uuid"
)

// GameUser представляет участника игры.
//
// @Description Модель участника игры с информацией об очках, звёздах и идентификаторах.
// @name GameUser
type GameUser struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	GameID    uuid.UUID `json:"game_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	GameName  string    `json:"game_name" example:"Math Battle"`
	GameCode  string    `json:"game_code" example:"ABC123"`
	FullName  string    `json:"full_name" example:"Иван Иванов"`
	Stars     int       `json:"stars" example:"3"`
	Score     int       `json:"score" example:"1200"`
	CreatedAt time.Time `json:"created_at" example:"2025-03-28T12:34:56Z"`
}
