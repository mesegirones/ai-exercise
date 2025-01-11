package rest

import (
	"context"
	"ia-exercise/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionService interface {
	GetQuestionList(ctx context.Context) ([]domain.Question, error)
}

type QuestionHandler struct {
	Service QuestionService
}

func NewQuestionHandler(r *gin.Engine, service QuestionService) {
	handler := &QuestionHandler{
		Service: service,
	}

	g := r.Group("/question")

	g.GET("/list", handler.GetQuestionList)

}

func (h *QuestionHandler) GetQuestionList(c *gin.Context) {
	response, err := h.Service.GetQuestionList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageResponse{Message: err.Error()})
	}
	c.JSON(http.StatusAccepted, response)
}
