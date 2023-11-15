package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ComapnyRepo interface {
	CreateComapny(company model.Company) (model.Company, error)
}

func NewCompanyRepo(db *gorm.DB) (ComapnyRepo, error) {
	if db == nil {
		log.Info().Msg("database cannot be nil")
		return nil, errors.New("database cannot be nil ")
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateComapny(company model.Company) (model.Company, error) {

	output := r.db.Create(&company)

	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in creating company table")
		return model.Company{}, errors.New("error in creating table")
	}

	return company, nil
}
