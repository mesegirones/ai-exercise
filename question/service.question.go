package question

import (
	"context"
	"fmt"
	"ia-exercise/domain"
	"strconv"
	"sync"
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

func (s *Service) ProcessQuestion(ctx context.Context, input *domain.QuestionInput) (domain.QuestionResponse, error) {

	response := domain.QuestionResponse{}
	response.Channel = make(chan domain.QuestionStatus)

	if input == nil {
		go s.InputError(response.Channel)
	}

	go s.AnalyzeQuestion(response.Channel, input)

	return response, nil
}

func (s *Service) InputError(channel chan domain.QuestionStatus) {
	channel <- domain.QuestionStatus{Status: "ERROR", Message: "missing input"}
	close(channel)
}

func (s *Service) AnalyzeQuestion(channel chan domain.QuestionStatus, input *domain.QuestionInput) {
	language, summary := s.FirstStep(channel, input.Question)
	response := s.SecondStep(channel, language, summary)

	channel <- domain.QuestionStatus{Status: "Answer obtained", Message: response}

	close(channel)
}

func (s *Service) FirstStep(channel chan domain.QuestionStatus, userInput string) (string, string) {
	results := make(chan string, 2)

	var wgFirstStep sync.WaitGroup

	wgFirstStep.Add(2)

	for i := 0; i < 2; i++ {
		status := "Invoking LLM " + strconv.Itoa(i+1)
		channel <- domain.QuestionStatus{Status: status, Message: ""}

		go s.LLM(&wgFirstStep, results, domain.LLMInput{
			Prompt:    "",
			UserInput: userInput,
		})

	}

	wgFirstStep.Wait()

	close(results)

	for result := range results {
		fmt.Println(result)
	}

	//TODO: handle results

	return "", ""
}

func (s *Service) SecondStep(channel chan domain.QuestionStatus, language string, summary string) string {
	results := make(chan string, 1)

	var wgSecondStep sync.WaitGroup
	wgSecondStep.Add(1)

	channel <- domain.QuestionStatus{Status: "Combining Results", Message: ""}

	go s.LLM(&wgSecondStep, results, domain.LLMInput{
		Prompt:    "",
		UserInput: "",
	})
	wgSecondStep.Wait()

	//TODO: handle results
	return ""
}

func (s *Service) LLM(wg *sync.WaitGroup, channel chan string, input domain.LLMInput) string {

	//TODO: handle IA stuf

	wg.Done()

	return ""
}
