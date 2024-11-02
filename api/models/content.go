package models

import (
	"content-app-backend/api/models/enums"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	//commonTags		string[]
	contentType		enums.ContentType
	//TODO 
}