package openai

import (
	"context"
	"ia-exercise/domain"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type ConfigOpenAI interface {
	GetAPIKEy() string
}

type LoggerProxy interface {
	Error(v ...interface{})
}

type Proxy struct {
	Client *openai.Client
	Logger LoggerProxy
}

func NewProxy(config ConfigOpenAI, logger LoggerProxy) (*Proxy, error) {
	client := openai.NewClient(
		option.WithAPIKey(config.GetAPIKEy()),
	)
	return &Proxy{
		Client: client,
		Logger: logger,
	}, nil
}

// OpenAI query with given promt.
func (p *Proxy) Query(ctx context.Context, message string) (string, error) {

	completion, err := p.Client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		p.Logger.Error(err)
		return "", domain.ErrServerError
	}

	return completion.Choices[0].Message.Content, nil
}
