package handler

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/middleware"
	"job-portal-api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupApi(auth authentication.Authenticaton, service service.Services) *gin.Engine {

	router := gin.New()

	mid, err := middleware.NewMid(auth)
	if err != nil {
		log.Panic("middleware are not set")
	}

	handler, err := NewHandler(service)
	if err != nil {
		log.Panic("handlers are not set")
	}

	router.Use(mid.Log(), gin.Recovery())

	router.POST("/api/Register", handler.)

}
