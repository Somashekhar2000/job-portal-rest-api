package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}
