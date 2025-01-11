package health

import (
	"context"
	"ia-exercise/domain"
	"time"
)

type Config interface {
	GetService() string
}

type HealthRepository interface {
	Check(ctx context.Context) error
}

type LoggerProxy interface {
	Error(v ...interface{})
}

type Service struct {
	ServiceName      string
	StartTime        time.Time
	HealthRepository HealthRepository
	Logger           LoggerProxy
}

func NewService(config Config, loggerProxy LoggerProxy) *Service {
	return &Service{
		ServiceName: config.GetService(),
		StartTime:   time.Now(),
		// HealthRepository: healthRepository,
		Logger: loggerProxy,
	}
}

func (s *Service) GetService() domain.ServiceResponse {
	return domain.ServiceResponse{
		Service: s.ServiceName,
		Message: "Up since " + s.StartTime.Format(time.RFC1123),
	}
}

func (s *Service) GetHealth(ctx context.Context) (domain.MessageResponse, error) {
	if err := s.HealthRepository.Check(ctx); err != nil {
		s.Logger.Error(err)
		return domain.MessageResponse{Message: "KO"}, domain.ErrInternalServerError
	}
	return domain.MessageResponse{Message: "OK"}, nil
}
