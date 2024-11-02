package repositories

import (
//	"content-app-backend/api/models"
	"gorm.io/gorm"
)

type IContentRepository interface {

}

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{db:db}
}