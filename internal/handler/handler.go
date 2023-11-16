package handler

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serviceUser    service.UserService
	serviceComapny service.ComapnyService
}

func SetupApi(auth authentication.Authenticaton, userService service.UserService, comapnyService service.ComapnyService) *gin.Engine {

	router := gin.New()

	mid, err := middleware.NewMid(auth)
	if err != nil {
		log.Panic("middleware are not set")
	}

	userHandler, err := NewUserHandler(userService)
	if err != nil {
		log.Panic("user handlers are not set")
	}

	companyHandler, err := NewCompanyHandler(comapnyService)
	if err != nil {
		log.Panic("company handlers are not set")
	}

	router.Use(mid.Log(), gin.Recovery())

	router.POST("/api/signup", userHandler.Signup)
	router.POST("/api/login", userHandler.login)

	router.POST("/api/create_comapny", companyHandler.AddCompany)
	router.GET("/api/get_company/:id", companyHandler.ViewCompanyByID)
	router.GET("/api/get_companies", companyHandler.ViewAllComapny)

	return router
}
