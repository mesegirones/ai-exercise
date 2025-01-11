package rest

import (
	"context"
	"io"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config interface {
	IsProd() bool
	IsLocal() bool
}

type LoggerProxy interface {
	GetInfoWriter() io.Writer
	GetErrorWriter() io.Writer
	Error(...interface{})
	Get() *slog.Logger
}

type OAuthProxy interface {
	AddAuthToContext(ctx context.Context, authHeader string) context.Context
}

func NewGinEngine(config Config, logger LoggerProxy) *gin.Engine {
	// Setting up Gin
	if config.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin.DefaultWriter = logger.GetInfoWriter()
	gin.DefaultErrorWriter = logger.GetErrorWriter()
	gin.DisableConsoleColor()

	r := gin.New()
	if config.IsLocal() {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(gin.Recovery())
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTION", "GET", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin", "User-Agent", "Accept-Encoding", "Accept", "Connection", "X-Api-Key"},
	}))

	return r
}
