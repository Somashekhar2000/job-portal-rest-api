package handler

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s service.UserService
}

func SetupApi(auth authentication.Authenticaton, service service.UserService) *gin.Engine {

	router := gin.New()

	mid, err := middleware.NewMid(auth)
	if err != nil {
		log.Panic("middleware are not set")
	}

	handler, err := NewUserHandler(service)
	if err != nil {
		log.Panic("handlers are not set")
	}

	router.Use(mid.Log(), gin.Recovery())

	router.POST("/api/signup", handler.Signup)
	router.POST("/api/login", handler.login)

}
