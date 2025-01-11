package rest

import (
	"github.com/gin-gonic/gin"
)

type EntityService interface {
}

type EntityHandler struct {
	Service EntityService
}

func NewEntityHandler(r *gin.Engine, svc EntityService) {
	_ = &EntityHandler{
		Service: svc,
	}
}
