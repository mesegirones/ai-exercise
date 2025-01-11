package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GetKeys() any
}

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(r *gin.Engine, svc AuthService) {
	handler := &AuthHandler{
		Service: svc,
	}

	g := r.Group("/api")
	g.GET("/auth/keys", handler.GetAuthKeys)
}

func (h *AuthHandler) GetAuthKeys(c *gin.Context) {
	c.JSON(http.StatusOK, h.Service.GetKeys())
}
