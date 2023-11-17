package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type JobRepository interface {
	CreateJob(jodData model.Job) (model.Response, error)
	ViewingJobByCompany(cID uint) ([]model.Job, error)
}

func NewJobRepo(db *gorm.DB) (JobRepository, error) {
	if db == nil {
		log.Info().Msg("database cannot be nil")
		return nil, errors.New("data base cannot be nil")
	}

	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateJob(jobData model.Job) (model.Response, error) {

	output := r.db.Create(&jobData)

	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in creating job table")
		return model.Response{}, errors.New("could not create job")
	}

	return model.Response{
		Id: jobData.ID,
	}, nil
}

func (r *Repo) ViewingJobByCompany(cID uint) ([]model.Job, error) {

	var jobData []model.Job

	output := r.db.Where("cid = ?", cID).Find(&jobData)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error ivalid company id")
		return nil, errors.New("invalid company id")
	}

	return jobData, nil
}
