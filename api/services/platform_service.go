package services

import (
//	"errors"
//	"content-app-backend/api/models"
	"content-app-backend/api/repositories"
)

type IPlatformService interface {

}

type PlatformService struct { 
	repo repositories.PlatformRepository
}

func NewPlatformService(repo repositories.PlatformRepository) IPlatformService {
	return &PlatformService{repo:repo}
}