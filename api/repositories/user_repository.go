package repositories

import (
//	"content-app-backend/api/models"
	"gorm.io/gorm"
)

type IUserRepository interface {

}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db:db}
}