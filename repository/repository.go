package repository

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

type Repository interface {
}

func NewRepository(db *gorm.DB) (Repository, error) {
	if db == nil {
		log.Info().Msg("database cannot be nil")
		return nil, errors.New("database cannot be nil")
	}
	return &Repo{
		db: db,
	}, nil
}
