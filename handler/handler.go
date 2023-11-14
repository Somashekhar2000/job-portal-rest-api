package handler

import (
	"errors"
	"job-portal-api/service"

	"github.com/rs/zerolog/log"
)

type Handler struct {
	service service.Services
}

type HandlerFuncs interface {
}

func NewHandler(service service.Services) (HandlerFuncs, error) {
	if service == nil {
		log.Info().Msg("service cannot be nil")
		return nil, errors.New("Service cannot be nil")
	}
	return &Handler{
		service: service,
	}, nil
}
