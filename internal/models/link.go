package models

import (
	"time"

	"github.com/google/uuid"
)

// GameLink представляет ссылку на игру, созданную пользователем.
//
// @Description Модель ссылки на игру с привязкой к пользователю, школой и классом.
// @name GameLink
type GameLink struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Code      string    `json:"code" example:"ABC123"`
	UserID    uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	GameName  string    `json:"game_name" example:"Исторический Квиз"`
	SchoolNum string    `json:"school_num" example:"42"`               // номер школы
	Class     string    `json:"class" example:"7А"`                    // например: "7А"
	Comment   string    `json:"comment" example:"Игра ко дню учителя"` // комментарий (опционально)
	CreatedAt time.Time `json:"created_at" example:"2025-03-28T14:00:00Z"`
}
