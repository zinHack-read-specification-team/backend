package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type DataRepository struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) *DataRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &DataRepository{db: db}
}

func (r *DataRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
