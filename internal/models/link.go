package models

import (
	"time"

	"github.com/google/uuid"
)

type GameLink struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	GameName  string    `json:"game_name" gorm:"not null"`
	SchoolNum string    `json:"school_num" gorm:"not null"` // номер школы
	Class     string    `json:"class" gorm:"not null"`      // например: "7А"
	Comment   string    `json:"comment" gorm:"type:text"`   // комментарий (опционально)
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
