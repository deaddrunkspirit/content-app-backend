package repositories

import (
//	"content-app-backend/api/models"
	"gorm.io/gorm"
)

type IPlatformRepository interface {

}

type PlatformRepository struct {
	db *gorm.DB
}

func NewPlatformRepository(db *gorm.DB) *PlatformRepository {
	return &PlatformRepository{db:db}
}