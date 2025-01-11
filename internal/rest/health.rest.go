package rest

import (
	"context"
	"ia-exercise/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthService interface {
	GetService() domain.ServiceResponse
	GetHealth(ctx context.Context) (domain.MessageResponse, error)
}

type HealthHandler struct {
	Service HealthService
}

func NewHealthHandler(r *gin.Engine, svc HealthService) {
	handler := &HealthHandler{
		Service: svc,
	}

	r.GET("/healthz", handler.GetHealth)
	r.GET("/", handler.GetService)
}

func (h *HealthHandler) GetService(c *gin.Context) {
	c.JSON(http.StatusOK, h.Service.GetService())
}

func (h *HealthHandler) GetHealth(c *gin.Context) {
	if r, err := h.Service.GetHealth(c); err != nil {
		c.JSON(http.StatusInternalServerError, r)
	} else {
		c.JSON(http.StatusOK, r)
	}
}
