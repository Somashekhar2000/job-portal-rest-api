package handler

import "github.com/gin-gonic/gin"

type CompanyHandler interface {
	AddCompany(c *gin.Context)
}

func (h *Handler) AddCompany(c *gin.Context) {

}
