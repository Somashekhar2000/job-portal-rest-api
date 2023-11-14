package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Register(c *gin.Context)
	Signin(c *gin.Context)
}

func (h *Handler) Register(c *gin.Context) {

}

func (h *Handler) Signin(c *gin.Context) {

}
