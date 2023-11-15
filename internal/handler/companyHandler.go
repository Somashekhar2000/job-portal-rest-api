package handler

import "github.com/gin-gonic/gin"

type CompanyHandler interface {
	AddCompany(c *gin.Context)
	ViewCompanyByID(c *gin.Context)
	ViewAllComapny(c *gin.Context)
}

func (h *Handler) AddCompany(c *gin.Context) {

}

func (h *Handler) ViewCompanyByID(c *gin.Context) {

}
func (h *Handler) ViewAllComapny(c *gin.Context) {

}
