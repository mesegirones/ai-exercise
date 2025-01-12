package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"ia-exercise/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionService interface {
	ProcessQuestion(ctx context.Context, input *domain.QuestionInput) (domain.QuestionResponse, error)
}

type QuestionHandler struct {
	Service QuestionService
}

func NewQuestionHandler(r *gin.Engine, service QuestionService) {
	handler := &QuestionHandler{
		Service: service,
	}

	g := r.Group("/question")

	g.POST("/ask", handler.AnalyzeQuestion)

}

func (h *QuestionHandler) AnalyzeQuestion(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	var input *domain.QuestionInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrBadParamInput)
		return
	}

	response, err := h.Service.ProcessQuestion(c, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageResponse{Message: err.Error()})
	}

	for status := range response.Channel {
		statusJSON, err := json.Marshal(status)

		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		}

		// Write status to the client using SSE format
		fmt.Fprintf(c.Writer, "data: %s\n\n", statusJSON)

		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush() // Flush the data to the client
		}
	}
}
