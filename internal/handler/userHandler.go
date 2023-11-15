package handler

import (
	"encoding/json"
	"errors"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/model"
	"job-portal-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type UserHandler interface {
	Signup(c *gin.Context)
	// Signin(c *gin.Context)
}

func NewUserHandler(service service.UserService) (UserHandler, error) {
	if service == nil {
		return nil, errors.New("userService Cannot be nil")
	}
	return &Handler{
		s: service,
	}, nil
}

func (h *Handler) Signup(c *gin.Context) {

	ctx := c.Request.Context()

	traceID, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing Trace Id in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userData model.UserSignup

	err := json.NewDecoder(c.Request.Body).Decode(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID : ", traceID).Msg("error in decoding signup struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID :", traceID).Msg("error in validating sigup struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	userdata, err := h.s.UserSignupDetail(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID :", traceID).Msg("error in user sigup")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, userdata)

}

func (h *Handler) Signin(c *gin.Context) {

}
