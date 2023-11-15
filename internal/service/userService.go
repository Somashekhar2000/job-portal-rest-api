package service

import (
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	passwordhash "job-portal-api/passwordHash"
)

type UserService interface {
	UserSignupDetail(userSignup model.UserSignup) (model.User, error)
}

func NewUserService(userRepo repository.UserRepository, a authentication.Authenticaton) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("User Repo cannot be nil")
	}
	return &Service{
		userRepo:       userRepo,
		authentication: a,
	}, nil
}

func (s *Service) UserSignupDetail(userData model.UserSignup) (model.User, error) {
	hashedPassword, err := passwordhash.HashingPassword(userData.Password)
	if err != nil {
		return model.User{}, err
	}

	userDetails := model.User{
		UserName: userData.UserName,
		EmailID:  userData.EmailID,
		Password: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return model.User{}, err
	}

	return userDetails, nil

}
