package main

import (
	"context"
	"fmt"
	"ia-exercise/health"
	"ia-exercise/internal/config"
	"ia-exercise/internal/proxy/logger"
	openai "ia-exercise/internal/proxy/openAI"
	"ia-exercise/internal/rest"
	"ia-exercise/question"
)

func main() {
	ctx := context.Background()

	// Setting up dependencies
	config := config.NewConfig()
	loggerProxy := logger.NewLoggerProxy(ctx, config.GetRestConfig())
	openaiProxy, err := openai.NewProxy(config.GetOpenAIConfig(), loggerProxy)
	if err != nil {
		loggerProxy.Error(err)
		return
	}

	//REST api config
	r := rest.NewGinEngine(config.GetRestConfig(), loggerProxy)

	healthService := health.NewService(config.GetHealthConfig(), loggerProxy)
	rest.NewHealthHandler(r, healthService)

	questionService := question.NewService(loggerProxy, openaiProxy)
	rest.NewQuestionHandler(r, questionService)

	if err := r.Run(fmt.Sprintf(":%s", config.GetRestConfig().GetPort())); err != nil {
		loggerProxy.Error(err)
	}

}
