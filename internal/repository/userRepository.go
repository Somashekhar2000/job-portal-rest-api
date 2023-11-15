package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(userData model.User) (model.User, error)
}

func NewUserRepo(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		return nil, errors.New("database cannot be nil")
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateUser(userData model.User) (model.User, error) {

	output := r.db.Create(&userData)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in database, could not create user")
	}

	return userData, nil
}