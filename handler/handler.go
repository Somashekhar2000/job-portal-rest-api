package handler

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/service"

	"github.com/gin-gonic/gin"
)

func SetupApi(auth authentication.Authenticaton, service service.Services) *gin.Engine {

}
