package question

import (
	"context"
	"ia-exercise/domain"
)

type LoggerProxy interface {
	Error(v ...interface{})
}

type Service struct {
	Logger LoggerProxy
}

func NewService(loggerProxy LoggerProxy) *Service {
	return &Service{
		Logger: loggerProxy,
	}
}

func (s *Service) GetQuestionList(ctx context.Context) ([]domain.Question, error) {

	return nil, nil
}
