package models

import (
    "gorm.io/gorm"
)

type User struct {
	gorm.Model
	login			string
	passwordHash	string
	//platforms		PlatformConnector
}