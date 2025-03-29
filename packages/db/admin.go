package db_package

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	username string
	password string
}
