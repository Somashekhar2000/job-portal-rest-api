package handler

import (
	"encoding/json"
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/model"
	"job-portal-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type CompanyHandler interface {
	AddCompany(c *gin.Context)
	ViewCompanyByID(c *gin.Context)
	ViewAllComapny(c *gin.Context)
}

func NewCompanyHandler(companyService service.ComapnyService) (CompanyHandler, error) {
	if companyService == nil {
		log.Info().Msg("error comapnyService cannot be nil")
		return nil, errors.New("error comapnyService cannot be nil")
	}

	return &Handler{
		serviceComapny: companyService,
	}, nil
}

func (h *Handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("trace Id missing")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}
	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("trace Id : ", traceId).Msg("login not success")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error ": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var companyData model.AddCompany

	err := json.NewDecoder(c.Request.Body).Decode(companyData)
	if err != nil {
		log.Error().Err(err).Str("trace Id : ", traceId).Msg("error in decpding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()

	err = validate.Struct(companyData)
	if err != nil {
		log.Error().Err(err).Msg("error in validating struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	company, err := h.serviceComapny.AddingCompany(companyData)
	if err != nil {
		log.Error().Err(err).Msg("error in creating company")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *Handler) ViewCompanyByID(c *gin.Context) {

}
func (h *Handler) ViewAllComapny(c *gin.Context) {

}
