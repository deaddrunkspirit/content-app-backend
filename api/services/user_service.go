package services

import (
//	"errors"
//	"content-app-backend/api/models"
	"content-app-backend/api/repositories"
)

type IUserService interface {

}

type UserService struct { 
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) IUserService {
	return &UserService{repo:repo}
}