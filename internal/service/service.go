package service

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/repository"
)

type Service struct {
	userRepo       repository.UserRepository
	authentication authentication.Authenticaton
}
