package models

import (
	"content-app-backend/api/models/enums"
	"gorm.io/gorm"
)

type PlatformConnector struct {
	gorm.Model
	name			string
	platformType	enums.PlatformType
	platformId 		int	
}