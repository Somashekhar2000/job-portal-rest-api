package service

import (
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/repository"

	"github.com/rs/zerolog/log"
)

type Service struct {
	repo           repository.Repository
	authentication authentication.Authenticaton
}

type Services interface {
}

func NewService(repo repository.Repository, authentication authentication.Authenticaton) (Services, error) {
	if repo == nil {
		log.Info().Msg("error repo4sitory interface in nil")
		return nil, errors.New("error repository interface cannot be nil")
	}
	return &Service{
		repo:           repo,
		authentication: authentication,
	}, nil
}
