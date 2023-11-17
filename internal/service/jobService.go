package service

import (
	"errors"
	"fmt"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type JobService interface {
	CreateJobByCompanyId(jobdata model.NewJobs, cID uint) (model.Response, error)
	ViewJobByCompanyID(cID uint) ([]model.Job, error)
	ViewJobByJobID(jID uint) (model.Job, error)
	ViewAllJobs() ([]model.Job, error)
}

func NewJobService(jobService repository.JobRepository) (JobService, error) {
	return &Service{
		jobRepo: jobService,
	}, nil
}

func (s *Service) CreateJobByCompanyId(jobDetails model.NewJobs, cID uint) (model.Response, error) {

	jobData := model.Job{
		Cid:             cID,
		Jobname:         jobDetails.Jobname,
		MinNoticePeriod: jobDetails.MinNoticePeriod,
		MaxNoticePeriod: jobDetails.MaxNoticePeriod,
		Description:     jobDetails.Description,
		Jobtype:         jobDetails.Jobtype,
		MinExperience:   jobDetails.MinExperience,
		MaxExperience:   jobDetails.MaxExperience,
	}

	for _, v := range jobDetails.Location {
		jobLocation := model.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Location = append(jobData.Location, jobLocation)
	}

	for _, v := range jobDetails.TechnologyStack {
		jobSkill := model.TechnologyStack{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.TechnologyStack = append(jobData.TechnologyStack, jobSkill)
	}

	for _, v := range jobDetails.Qualifications {
		jobQualification := model.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Qualifications = append(jobData.Qualifications, jobQualification)
	}

	for _, v := range jobDetails.Shift {
		jobShift := model.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Shift = append(jobData.Shift, jobShift)
	}

	responseData, err := s.jobRepo.CreateJob(jobData)
	if err != nil {
		return model.Response{}, err
	}

	return responseData, nil

}

func (s *Service) ViewJobByCompanyID(cID uint) ([]model.Job, error) {

	jobData, err := s.jobRepo.GetJobByCompanyID(cID)

	if err != nil {
		return nil, err
	}

	if jobData == nil {
		log.Error().Err(errors.New("error jobs does not exists in this company"))
		return nil, err
	}

	return jobData, nil
}

func (s *Service) ViewJobByJobID(jID uint) (model.Job, error) {

	jobData, err := s.jobRepo.GetJobByJobID(jID)
	if err != nil {
		fmt.Println("========-------==========", err)
		return model.Job{}, err
	}
	fmt.Println("[[[[[[[[[[[[]]]]]]]]]]]]")
	return jobData, nil
}

func (s *Service) ViewAllJobs() ([]model.Job, error) {
	jobData, err := s.jobRepo.GetAllJobs()
	if err != nil {
		return nil, err
	}

	return jobData, nil
}
