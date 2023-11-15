package handler

import "github.com/gin-gonic/gin"

type JobHandler interface {
	CreateJobByCompanyID(c *gin.Context)
	ViewJobByCompanyId(c *gin.Context)
	ViewJobByJobID(c *gin.Context)
	ViewAllJobs(c *gin.Context)
	ProcessJobApplication(c *gin.Context)
}

func (h *Handler) CreateJob(c *gin.Context) {

}

func (h *Handler) ViewJobByCompanyId(c *gin.Context) {

}

func (h *Handler) ViewJobByJobID(c *gin.Context) {

}

func (h *Handler) ViewAllJobs(c *gin.Context) {

}

func (h *Handler) ProcessJobApplication(c *gin.Context) {

}
