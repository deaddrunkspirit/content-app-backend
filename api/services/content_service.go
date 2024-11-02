package services

import (
//	"errors"
//	"content-app-backend/api/models"
	"content-app-backend/api/repositories"
)

type IContentService interface {

}

type ContentService struct { 
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) IContentService {
	return &ContentService{repo:repo}
}