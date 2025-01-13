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
	Debug(v ...interface{})
}

type OpenaiProxy interface {
	Query(ctx context.Context, message string) (string, error)
}

type Service struct {
	Logger      LoggerProxy
	OpenaiProxy OpenaiProxy
}

func NewService(loggerProxy LoggerProxy, openaiProxy OpenaiProxy) *Service {
	return &Service{
		Logger:      loggerProxy,
		OpenaiProxy: openaiProxy,
	}
}

// Validating input and generating analysis
func (s *Service) ProcessQuestion(ctx context.Context, input *domain.QuestionInput) (domain.QuestionResponse, error) {

	response := domain.QuestionResponse{}
	response.Channel = make(chan domain.QuestionStatus)

	if input == nil {
		go s.InputError(response.Channel)
	}

	go s.AnalyzeQuestion(ctx, response.Channel, input)

	return response, nil
}

func (s *Service) InputError(channel chan domain.QuestionStatus) {
	channel <- domain.QuestionStatus{Status: "ERROR", Message: "missing input"}
	close(channel)
}

// Main analysis query
func (s *Service) AnalyzeQuestion(ctx context.Context, channel chan domain.QuestionStatus, input *domain.QuestionInput) {
	language, summary := s.FirstStep(ctx, channel, input.Question)
	response := s.SecondStep(ctx, channel, language, summary)

	channel <- domain.QuestionStatus{Status: "Answer obtained", Message: response}

	close(channel)
}

// First step of the analysis. Calls de first two LLM and returns the results in string format.
func (s *Service) FirstStep(ctx context.Context, channel chan domain.QuestionStatus, userInput string) (string, string) {
	results := make(chan domain.LLMOutput, 2)

	var wgFirstStep sync.WaitGroup

	wgFirstStep.Add(2)

	prompts := []domain.LLMInput{
		{
			Prompt:   "In one word, wich is the main point of the following text: %s",
			DataType: domain.DataTypeEnumSUMMARY,
		},
		{
			Prompt:   "In maximum two words, in which langage and slang is written the follwing text: %s",
			DataType: domain.DataTypeEnumLANGUAGE,
		},
	}

	for i, input := range prompts {
		status := "Invoking LLM " + strconv.Itoa(i+1)
		channel <- domain.QuestionStatus{Status: status, Message: ""}
		go s.LLM(ctx, &wgFirstStep, results, fmt.Sprintf(input.Prompt, userInput), input.DataType)

	}

	wgFirstStep.Wait()

	close(results)

	var language, summary string
	for result := range results {
		switch result.DataType {
		case domain.DataTypeEnumLANGUAGE:
			language = result.Message
		case domain.DataTypeEnumSUMMARY:
			summary = result.Message
		}
	}

	return language, summary
}

// Calls the third LLM and returns result in string format.
func (s *Service) SecondStep(ctx context.Context, channel chan domain.QuestionStatus, language string, summary string) string {
	results := make(chan domain.LLMOutput, 1)

	var wgSecondStep sync.WaitGroup
	wgSecondStep.Add(1)

	channel <- domain.QuestionStatus{Status: "Combining Results", Message: ""}

	promt := "Give me a short fun fact using maxomum 20 words about this topic %s witten in %s"
	go s.LLM(ctx, &wgSecondStep, results, fmt.Sprintf(promt, language, summary), domain.DataTypeEnumRESULT)
	wgSecondStep.Wait()

	close(results)

	var response string
	for result := range results {
		if result.DataType == domain.DataTypeEnumRESULT {
			response = result.Message
		}
	}
	return response
}

// Perform LLM query
func (s *Service) LLM(ctx context.Context, wg *sync.WaitGroup, results chan domain.LLMOutput, promt string, dataType domain.DataTypeEnum) {
	defer wg.Done()

	response, err := s.OpenaiProxy.Query(ctx, promt)

	if err != nil {
		results <- domain.LLMOutput{
			DataType: domain.DataTypeEnumERROR,
			Message:  "ERROR: Somthing wrong happened :(",
		}
	}
	results <- domain.LLMOutput{
		DataType: dataType,
		Message:  response,
	}

}
